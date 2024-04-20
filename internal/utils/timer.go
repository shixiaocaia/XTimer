package utils

import (
	"fmt"
	"time"

	"github.com/gorhill/cronexpr"
)

func GetEnableLockKey(app string) string {
	return fmt.Sprintf("enable_timer_lock_%s", app)
}

func GetTimeBatch(cron string, end time.Time) ([]time.Time, error) {
	start := time.Now()
	if end.Before(start) {
		return nil, fmt.Errorf("end can not earlier than start, start: %v, end: %v", start, end)
	}

	// 解析这个时间
	expr, err := cronexpr.Parse(cron)
	if err != nil {
		return nil, err
	}

	// 基于一个时间步的时间
	var nexts []time.Time
	for start.Before(end) {
		next := expr.Next(start)
		if next.UnixNano() < 0 {
			return nil, fmt.Errorf("fail to parse time from cron: %s", cron)
		}
		nexts = append(nexts, next)
		start = next
	}

	return nexts, nil
}
