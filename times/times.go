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
func GetYearMonth() string {
	return time.Now().Format("200601")
}
