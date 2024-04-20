package constant

type TimerStatus int

func (t TimerStatus) ToInt() int {
	return int(t)
}

const (
	Unabled TimerStatus = 1
	Enabled TimerStatus = 2
)
