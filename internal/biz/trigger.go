package biz

import (
	"Xtimer/internal/conf"
	"Xtimer/internal/constant"
	"Xtimer/internal/utils"
	"Xtimer/third_party/log"
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

type TriggerUseCase struct {
	confData  *conf.Data
	timerRepo TimerRepo
	taskRepo  TimerTaskRepo
	taskCache TaskCacheRepo
	tm        Transaction
	pool      WorkerPool

	executor *ExecutorUseCase
}

func NewTriggerUseCase(confData *conf.Data, timerRepo TimerRepo, taskRepo TimerTaskRepo, tm Transaction, taskCache TaskCacheRepo, executor *ExecutorUseCase) *TriggerUseCase {
	return &TriggerUseCase{
		confData:  confData,
		timerRepo: timerRepo,
		taskRepo:  taskRepo,
		taskCache: taskCache,
		tm:        tm,
		pool:      NewGoWorkerPool(int(confData.Trigger.WorkersNum)),
		executor:  executor,
	}
}

func (w *TriggerUseCase) Work(ctx context.Context, minuteBucketKey string, ack func()) error {
	startTime, err := utils.GetStartMinute(minuteBucketKey)
	if err != nil {
		return err
	}

	ticker := time.NewTicker(time.Duration(w.confData.Trigger.ZrangeGapSeconds) * time.Second)
	defer ticker.Stop()

	notifier := NewSafeChan(int(time.Minute / (time.Duration(w.confData.Trigger.ZrangeGapSeconds) * time.Second)))
	defer notifier.Close()

	endTime := startTime.Add(time.Minute)
	var wg sync.WaitGroup
	for range ticker.C {
		select {
		case e := <-notifier.GetChan():
			err, _ = e.(error)
			return err
		default:
		}

		wg.Add(1)
		nextTime := startTime.Add(time.Duration(w.confData.Trigger.ZrangeGapSeconds) * time.Second)
		go func(startTime, nextTime time.Time) {
			defer wg.Done()
			if err := w.handleBatch(ctx, minuteBucketKey, startTime, nextTime); err != nil {
				notifier.Put(err)
			}
		}(startTime, nextTime)

		startTime = startTime.Add(time.Duration(w.confData.Trigger.ZrangeGapSeconds) * time.Second)
		if startTime.Equal(endTime) || startTime.After(endTime) {
			break
		}
	}

	wg.Wait()
	select {
	case e := <-notifier.GetChan():
		err, _ = e.(error)
		return err
	default:
	}

	// 成功执行完任务，延迟分布式锁
	ack()
	log.InfoContextf(ctx, "ack success, key: %s", minuteBucketKey)

	return nil
}

func (w *TriggerUseCase) handleBatch(ctx context.Context, key string, start, end time.Time) error {
	bucket, err := getBucket(key)
	if err != nil {
		return err
	}

	tasks, err := w.getTasksByTime(ctx, key, bucket, start, end)
	if err != nil {
		return err
	}

	// 没有任务需要执行
	if len(tasks) == 0 {
		log.DebugContextf(ctx, "no task to execute, key: %s", key)
		return nil
	}
	log.DebugContextf(ctx, "tasks to execute, key: %s, tasks: %v", key, tasks)

	//timerIds := make([]int64, 0, len(tasks))
	//for _, task := range tasks {
	//	timerIds = append(timerIds, task.TimerID)
	//}

	for _, task := range tasks {
		if err := w.pool.Submit(func() {
			if err := w.executor.Work(ctx, utils.UnionTimerIDUnix(uint(task.TimerID), task.RunTimer)); err != nil {
				log.ErrorContextf(ctx, "executor work-%v failed, err: %v", utils.UnionTimerIDUnix(uint(task.TimerID), task.RunTimer), err)
			}
		}); err != nil {
			return err
		}
	}
	return nil
}

func getBucket(slice string) (int, error) {
	timeBucket := strings.Split(slice, "_")
	if len(timeBucket) != 2 {
		return -1, fmt.Errorf("invalid format of msg key: %s", slice)
	}
	return strconv.Atoi(timeBucket[1])
}

func (w *TriggerUseCase) getTasksByTime(ctx context.Context, key string, bucket int, start, end time.Time) ([]*TimerTask, error) {
	// 先走缓存
	tasks, err := w.taskCache.GetTasksByTime(ctx, key, start.UnixMilli(), end.UnixMilli())
	if err == nil || len(tasks) == 0 {
		return tasks, nil
	}

	// 倘若缓存查询报错，再走db
	tasks, err = w.taskRepo.GetTasksByTimeRange(ctx, start.UnixMilli(), end.UnixMilli(), constant.TaskStatusNotRun.ToInt())
	if err != nil {
		return nil, err
	}

	maxBucket := w.confData.Scheduler.BucketsNum
	var validTask []*TimerTask
	for _, task := range tasks {
		if uint(task.TimerID)%uint(maxBucket) != uint(bucket) {
			continue
		}
		validTask = append(validTask, task)
	}

	return validTask, nil
}
