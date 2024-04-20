package utils

import "fmt"

func GetEnableLockKey(app string) string {
	return fmt.Sprintf("enable_timer_lock_%s", app)
}
