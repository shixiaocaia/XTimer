package utils

import (
	"Xtimer/internal/constant"
	"fmt"
	cronexpr "github.com/robfig/cron/v3"
	"strconv"
	"strings"
	"time"
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

func GetStartMinute(str string) (time.Time, error) {
	// 2024-04-15 13:38_2
	timeBucket := strings.Split(str, "_")
	if len(timeBucket) != 2 {
		return time.Time{}, fmt.Errorf("invalid format of msg key: %s", str)
	}

	return time.ParseInLocation(constant.MinuteFormat, timeBucket[0], time.Local)
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

func GetSliceMsgKey(t time.Time, bucketID int) string {
	return fmt.Sprintf("%s_%d", t.Format(constant.MinuteFormat), bucketID)
}

func GetTimeBatch(cron string, end time.Time) ([]time.Time, error) {
	start := time.Now()
	if end.Before(start) {
		return nil, fmt.Errorf("end can not earlier than start, start: %v, end: %v", start, end)
	}

	specParser := cronexpr.NewParser(cronexpr.Second | cronexpr.Minute | cronexpr.Hour | cronexpr.Dom | cronexpr.Month | cronexpr.Dow)
	sched, err := specParser.Parse(cron)
	if err != nil {
		return nil, err
	}

	var next []time.Time
	for start.Before(end) {
		n := sched.Next(start)
		if n.UnixNano() < 0 {
			return nil, fmt.Errorf("fail to parse time from cron: %s", cron)
		}
		next = append(next, n)
		start = n
	}

	return next, nil
}
