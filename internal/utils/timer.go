package utils

import (
	"Xtimer/internal/constant"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gorhill/cronexpr"
)

func GetEnableLockKey(app string) string {
	return fmt.Sprintf("enable_timer_lock_%s", app)
}

func UnionTimerIDUnix(timeID uint, unix int64) string {
	return fmt.Sprintf("%d_%d", timeID, unix)
}

func GetTimeBucketLockKey(t time.Time, bucketID int) string {
	return fmt.Sprintf("time_bucket_lock_%s_%d", t.Format(constant.MinuteFormat), bucketID)
}

func SplitTimerIDUnix(str string) (int64, int64, error) {
	timerIDUnix := strings.Split(str, "_")
	if len(timerIDUnix) != 2 {
		return 0, 0, fmt.Errorf("invalid timerID unix str: %s", str)
	}

	timerID, _ := strconv.ParseInt(timerIDUnix[0], 10, 64)
	unix, _ := strconv.ParseInt(timerIDUnix[1], 10, 64)
	return timerID, unix, nil
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
