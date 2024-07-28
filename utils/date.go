package utils

import "time"

func GetCurrentDate() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02")
}
