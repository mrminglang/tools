package times

import (
	"github.com/jinzhu/now"
	"time"
)

const LocalFormat = "2006-01-02 15:04:05"

// StartOfDay 当天开始时间
func StartOfDay() string {
	return now.BeginningOfDay().Format(LocalFormat)
}

// EndOfDay 当天结束时间
func EndOfDay() string {
	return now.EndOfDay().Format(LocalFormat)
}

// StartOfMonth 当月开始时间
func StartOfMonth() string {
	return now.BeginningOfMonth().Format(LocalFormat)
}

// EndOfMonth 当月结束时间
func EndOfMonth() string {
	return now.EndOfMonth().Format(LocalFormat)
}

// StartOfYear 当年开始时间
func StartOfYear() string {
	return now.BeginningOfYear().Format(LocalFormat)
}

// EndOfYear 当年结束时间
func EndOfYear() string {
	return now.EndOfYear().Format(LocalFormat)
}

// 获取当前时间的年月份值
func GetYearMonth(layout string) string {
	return time.Now().Format(layout)
}

// 获取当前时间下月的年月份值
func GetYearProximo(layout string) string {
	currentTime := time.Now()
	nextMonth := currentTime.AddDate(0, 1, 0)
	nextMonthValue := nextMonth.Format(layout)
	return nextMonthValue
}

// 获取开始时间如: 20230418 到2023-04-18 00:00:00
func GetStartTime(timeStr string, layout string) string {
	startTime, _ := time.Parse(layout, timeStr)
	return startTime.Format(LocalFormat)
}

// 获取结束时间如: 20230418 到2023-04-18 23:59:59
func GetEndtTime(timeStr string, layout string) string {
	startTime, _ := time.Parse(layout, timeStr)
	endTime := startTime.Add(time.Hour*24 - time.Second)
	return endTime.Format(LocalFormat)
}
