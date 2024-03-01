package utime

import "time"

func ZeroHour(day int) time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day()+day, 0, 0, 0, 0, now.Location())
}
