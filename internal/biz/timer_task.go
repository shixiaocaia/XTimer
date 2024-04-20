package biz

import "context" // 表定义也需要放在biz层, 为了解耦biz与data层

type TimerTask struct {
	TaskID   int64  `gorm:"column:task_id"`
	App      string `gorm:"column:app;NOT NULL"`           // 定义ID
	TimerID  int64  `gorm:"column:timer_id;NOT NULL"`      // 定义ID
	Output   string `gorm:"column:output;default:null"`    // 执行结果
	RunTimer int64  `gorm:"column:run_timer;default:null"` // 执行时间
	CostTime int    `gorm:"column:cost_time"`              // 执行耗时
	Status   int    `gorm:"column:status;NOT NULL"`        // 当前状态
}

func (t *TimerTask) TableName() string {
	return "timer_task"
}

type TimerTaskRepo interface {
	BatchSave(context.Context, []*TimerTask) error
	Update(context.Context, *TimerTask) (*TimerTask, error)
	GetTasksByTimeRange(context.Context, int64, int64, int) ([]*TimerTask, error)
	GetTasksByTimerIdAndRunTimer(context.Context, int64, int64) (*TimerTask, error)
}

type TaskCache interface {
	BatchCreateTasks(ctx context.Context, tasks []*TimerTask) error
	GetTasksByTime(ctx context.Context, table string, start, end int64) ([]*TimerTask, error)
}
