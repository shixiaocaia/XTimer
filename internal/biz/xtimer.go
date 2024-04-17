package biz

import (
	"Xtimer/internal/conf"
	"context"
	"time"
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

type XTimerRepo interface {
	Save(context.Context, *Timer) (*Timer, error)
	Update(context.Context, *Timer) (*Timer, error)
	FindByID(context.Context, int64) (*Timer, error)
	FindByStatus(context.Context, int) ([]*Timer, error)
	Delete(context.Context, int64) error
}

type XTimerUseCase struct {
	confData  *conf.Data
	timerRepo XTimerRepo
}

func NewXTimerUseCase(confData *conf.Data, timerRepo XTimerRepo) *XTimerUseCase {
	return &XTimerUseCase{
		confData:  confData,
		timerRepo: timerRepo,
	}
}

func (uc *XTimerUseCase) CreateTimer(ctx context.Context, g *Timer) (*Timer, error) {
	return uc.timerRepo.Save(ctx, g)
}
