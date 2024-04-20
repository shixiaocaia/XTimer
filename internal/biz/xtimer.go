package biz

import (
	"Xtimer/internal/conf"
	"Xtimer/internal/constant"
	"Xtimer/internal/utils"
	"Xtimer/third_party/lock"
	"Xtimer/third_party/log"
	"context"
	"github.com/pkg/errors"
	context2 "golang.org/x/net/context"
	"time"
)

const (
	defaultEnableGapSeconds = 3
)

// 表定义也需要放在biz层, 还是为了解耦biz与data层
type Timer struct {
	TimerId         int64 `gorm:"column:id"`
	App             string
	Name            string
	Status          int
	Cron            string
	NotifyHTTPParam string     `gorm:"column:notify_http_param;NOT NULL" json:"notify_http_param,omitempty"` // Http 回调参数
	CreateTime      *time.Time `gorm:"column:create_time;default:null"`
	ModifyTime      *time.Time `gorm:"column:modify_time;default:null"`
}

// TableName 表名
func (p *Timer) TableName() string {
	return "xtimer"
}

type TimerRepo interface {
	Save(context.Context, *Timer) (*Timer, error)
	Update(context.Context, *Timer) (*Timer, error)
	FindByID(context.Context, int64) (*Timer, error)
	FindByStatus(context.Context, int) ([]*Timer, error)
	Delete(context.Context, int64) error
}

type XTimerUseCase struct {
	confData      *conf.Data
	timerRepo     TimerRepo
	taskRepo      TimerTaskRepo
	taskCacheRepo TaskCacheRepo
	tm            Transaction

	muc *MigratorUseCase
}

func NewXTimerUseCase(confData *conf.Data, timerRepo TimerRepo, tm Transaction, taskRepo TimerTaskRepo, muc *MigratorUseCase, taskCache TaskCacheRepo) *XTimerUseCase {
	return &XTimerUseCase{
		confData:      confData,
		timerRepo:     timerRepo,
		tm:            tm,
		taskRepo:      taskRepo,
		taskCacheRepo: taskCache,
		muc:           muc,
	}
}

func (uc *XTimerUseCase) CreateTimer(ctx context.Context, g *Timer) (*Timer, error) {
	return uc.timerRepo.Save(ctx, g)
}

func (uc *XTimerUseCase) ActiveTimer(ctx context.Context, app string, timerId int64, status int32) error {
	// 限制激活和去激活频次
	locker := lock.NewRedisLock(utils.GetEnableLockKey(app),
		lock.WithExpireSeconds(defaultEnableGapSeconds),
		lock.WithWatchDogMode())
	defer func(locker *lock.RedisLock, ctx context2.Context) {
		err := locker.Unlock(ctx)
		if err != nil {
			log.ErrorContextf(ctx, "EnableTimer 自动解锁失败", err.Error())
		}
	}(locker, ctx)
	err := locker.Lock(ctx)
	// 抢锁失败, 直接跳过执行, 下一轮
	if err != nil {
		log.InfoContextf(ctx, "激活/去激活操作过于频繁，请稍后再试！", err.Error())
		return errors.New("激活/去激活操作过于频繁，请稍后再试！")
	}

	uc.tm.InTx(ctx, func(ctx context.Context) error {
		// 1. 数据库获取Timer
		timer, err := uc.timerRepo.FindByID(ctx, timerId)
		if err != nil {
			log.ErrorContextf(ctx, "ActiveTimer failed:, timer not found, err: %v", err.Error())
			return err
		}

		// 2. 校验状态
		if timer.Status == int(status) {
			log.InfoContextf(ctx, "ActiveTimer failed: status is the same, timerId: %v, status: %v", timerId, status)
			return nil
		}

		// 3. 更新status
		timer.Status = int(status)
		if _, err := uc.timerRepo.Update(ctx, timer); err != nil {
			log.ErrorContextf(ctx, "ActiveTimer failed: update status failed, timerId: %v, status: %v, err: %v", timerId, status, err)
			return err
		}

		// 4. 如果是激活状态, 生成一批任务
		if timer.Status == constant.Enabled.ToInt() {
			if err := uc.muc.MigratorTimer(ctx, timer); err != nil {
				log.ErrorContextf(ctx, "ActiveTimer failed: MigratorTimer failed, timerId: %v, err: %v", timerId, err)
				return err
			}
		}
		return nil
	})

	return nil
}
