package constant

type TimerStatus int
type TaskStatus int

func (t TimerStatus) ToInt() int {
	return int(t)
}

const (
	Enabled  TimerStatus = 1
	Disabled TimerStatus = 2
)

const (
	TaskStatusNotRun  TaskStatus = 0
	TaskStatusSuccess TaskStatus = 1
	TaskStatusFail    TaskStatus = 2
)

func (t TaskStatus) ToInt() int {
	return int(t)
}

const (
	MinuteFormat = "2006-01-02 15:04"
)
