package biz

import (
	"Xtimer/internal/conf"
	"Xtimer/internal/constant"
	"Xtimer/internal/utils"
	"Xtimer/third_party/log"
	"context"
	"fmt"
	"time"
)

type MigratorUseCase struct {
	confData      *conf.Data
	timerRepo     TimerRepo
	taskRepo      TimerTaskRepo
	taskCacheRepo TaskCacheRepo
}

func NewMigratorUseCase(confData *conf.Data, timerRepo TimerRepo, taskRepo TimerTaskRepo, taskCache TaskCacheRepo) *MigratorUseCase {
	return &MigratorUseCase{
		confData:      confData,
		timerRepo:     timerRepo,
		taskRepo:      taskRepo,
		taskCacheRepo: taskCache,
	}
}

func (uc *MigratorUseCase) BatchMigratorTimer(ctx context.Context) error {
	timers, err := uc.timerRepo.FindByStatus(ctx, constant.Enabled.ToInt())
	if err != nil {
		log.ErrorContextf(ctx, "批量迁移Timer失败, 获取定时器失败, err: %v", err)
		return err
	}
	for _, timer := range timers {
		// todo 加锁，分布式情况下避免重复任务生成
		err = uc.MigratorTimer(ctx, timer)
		if err != nil {
			log.ErrorContextf(ctx, "批量迁移, 迁移单个Timer失败, timerId:%v, err:%v", timer.TimerId, err)
		}
		time.Sleep(5 * time.Second)
	}
	return nil
}

func (uc *MigratorUseCase) MigratorTimer(ctx context.Context, timer *Timer) error {
	// 1. 校验状态, 只有Enable状态的Timer才能迁移
	if timer.Status != constant.Enabled.ToInt() {
		return fmt.Errorf("Timer非Unable状态，迁移失败，timerId:: %d", timer.TimerId)
	}

	// 2. 取得批量的执行时机
	start := time.Now()
	// [start, start + 2 * 30min]
	end := start.Add(2 * time.Duration(uc.confData.GetMigrator().MigrateStepMinutes) * time.Minute)
	executeTimes, err := utils.GetTimeBatch(timer.Cron, end)
	if err != nil {
		log.ErrorContextf(ctx, "get executeTimes failed, err: %v", err)
		return err
	}
	//if len(executeTimes) == 0 {
	//	return nil
	//}

	// 3. 创建任务, 插入MySQL
	tasks := timer.BatchTasksFromTimer(executeTimes)
	if err := uc.taskRepo.BatchSave(ctx, tasks); err != nil {
		log.ErrorContextf(ctx, "DB存储tasks失败: %v", err)
		return err
	}
	// 4. 插入Redis
	if err := uc.taskCacheRepo.BatchCreateTasks(ctx, tasks); err != nil {
		log.ErrorContextf(ctx, "Redis存储tasks失败: %v", err)
		return err
	}

	return nil
}

func (t *Timer) BatchTasksFromTimer(executeTimes []time.Time) []*TimerTask {
	tasks := make([]*TimerTask, 0, len(executeTimes))
	for _, executeTime := range executeTimes {
		tasks = append(tasks, &TimerTask{
			App:      t.App,
			TimerID:  t.TimerId,
			Status:   constant.TaskStatusNotRun.ToInt(),
			RunTimer: executeTime.UnixMilli(),
		})
	}
	return tasks
}
