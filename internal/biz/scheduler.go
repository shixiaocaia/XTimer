package biz

import (
	"Xtimer/internal/conf"
	"Xtimer/internal/utils"
	"Xtimer/third_party/lock"
	"Xtimer/third_party/log"
	"context"
	"time"
)

type SchedulerUseCase struct {
	confData  *conf.Data
	timerRepo TimerRepo
	taskRepo  TimerTaskRepo
	taskCache TaskCacheRepo
	tm        Transaction
	pool      WorkerPool

	trigger *TriggerUseCase
}

func NewSchedulerUseCase(confData *conf.Data, timerRepo TimerRepo, taskRepo TimerTaskRepo, taskCache TaskCacheRepo, tm Transaction, trigger *TriggerUseCase) *SchedulerUseCase {
	return &SchedulerUseCase{
		confData:  confData,
		timerRepo: timerRepo,
		taskRepo:  taskRepo,
		taskCache: taskCache,
		pool:      NewGoWorkerPool(int(confData.Scheduler.WorkersNum)),
		tm:        tm,
		trigger:   trigger,
	}
}

func (w *SchedulerUseCase) Work(ctx context.Context) error {
	ticker := time.NewTicker(time.Duration(w.confData.Scheduler.TryLockGapMilliSeconds) * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		select {
		case <-ctx.Done():
			log.WarnContextf(ctx, "stopped")
			return nil
		default:
		}

		w.handleSlices(ctx)
	}
	return nil
}

func (w *SchedulerUseCase) handleSlices(ctx context.Context) {
	bucket := int(w.confData.Scheduler.BucketsNum)
	for i := 0; i < bucket; i++ {
		w.handleSlice(ctx, i)
	}
}

func (w *SchedulerUseCase) handleSlice(ctx context.Context, bucketID int) {
	defer func() {
		// 捕获异常，不再上报
		if r := recover(); r != nil {
			log.ErrorContextf(ctx, "handleSlice %v run err. Recovered from panic:%v", bucketID, r)
		}
	}()

	//log.InfoContextf(ctx, "scheduler_%v start: %v", bucketID, time.Now())

	now := time.Now()
	// 如果能获取到上一分钟的锁，说明上一分钟任务没有全部处理完成，没有延时锁
	// 重新执行上一分钟任务
	if err := w.pool.Submit(func() {
		w.asyncHandleSlice(ctx, now.Add(-time.Minute), bucketID)
	}); err != nil {
		log.ErrorContextf(ctx, "[handle slice] submit task failed, err: %v", err)
	}
	if err := w.pool.Submit(func() {
		w.asyncHandleSlice(ctx, now, bucketID)
	}); err != nil {
		log.ErrorContextf(ctx, "[handle slice] submit task failed, err: %v", err)
	}

	//log.InfoContextf(ctx, "scheduler_%v end: %v", bucketID, time.Now())
}

func (w *SchedulerUseCase) asyncHandleSlice(ctx context.Context, t time.Time, bucketID int) {
	// 限制激活和去激活频次
	locker := lock.NewRedisLock(utils.GetTimeBucketLockKey(t, bucketID),
		lock.WithExpireSeconds(w.confData.Scheduler.TryLockSeconds))
	err := locker.Lock(ctx)
	if err != nil {
		// log.InfoContextf(ctx, "asyncHandleSlice 获取分布式锁失败: %v", err.Error())
		// 抢锁失败, 直接跳过执行, 下一轮
		return
	}
	// 保证每个分钟时间桶，只有一个协程处理
	log.InfoContextf(ctx, "get scheduler lock success, key: %s", utils.GetTimeBucketLockKey(t, bucketID))

	// 成功后延锁，避免下一个时钟分片获取到重复执行
	ack := func() {
		if err := locker.DelayExpire(ctx, w.confData.Scheduler.SuccessExpireSeconds); err != nil {
			log.ErrorContextf(ctx, "expire lock failed, lock key: %s, err: %v", utils.GetTimeBucketLockKey(t, bucketID), err)
		} else {
			log.DebugContextf(ctx, "expire lock success, lock key: %s", utils.GetTimeBucketLockKey(t, bucketID))
		}
	}

	if err := w.trigger.Work(ctx, utils.GetSliceMsgKey(t, bucketID), ack); err != nil {
		log.ErrorContextf(ctx, "trigger work failed, SliceMsgKey[%v] err: %v", utils.GetSliceMsgKey(t, bucketID), err)
	}

}
