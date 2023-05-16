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

// 获取指定时间到当前时间的年月IDs,如202303 返回["202303","202304"]
func GenerateYearMonthIds(timeStr string, layout string) []string {
	var result []string
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		return result
	}
	for t.Before(time.Now()) {
		result = append(result, t.Format(layout))
		t = t.AddDate(0, 1, 0)
	}
	return result
}

// 获取昨天的时间
func GetYesterday(layout string) string {
	// 获取当前时间
	currentTime := time.Now()
	// 前一天时间
	yesterday := currentTime.AddDate(0, 0, -1)
	// 格式化时间
	return yesterday.Format(layout)
}

// 获取指定时区的时间
func GetTimeZone(currentTime string, timeZone, layout string) string {
	parseTime, err := time.Parse(LocalFormat, currentTime)
	if err != nil {
		return currentTime
	}

	locationTime, err := time.LoadLocation(timeZone)
	if err != nil {
		return currentTime
	}

	return parseTime.In(locationTime).Format(layout)
}
