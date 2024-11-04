package times

import (
	"errors"
	"fmt"
	"github.com/jinzhu/now"
	"strconv"
	"time"
)

// StartOfDay 当天开始时间
func StartOfDay() string {
	return now.BeginningOfDay().Format(TimeFormat)
}

// EndOfDay 当天结束时间
func EndOfDay() string {
	return now.EndOfDay().Format(TimeFormat)
}

// StartOfMonth 当月开始时间
func StartOfMonth() string {
	return now.BeginningOfMonth().Format(TimeFormat)
}

// EndOfMonth 当月结束时间
func EndOfMonth() string {
	return now.EndOfMonth().Format(TimeFormat)
}

// StartOfYear 当年开始时间
func StartOfYear() string {
	return now.BeginningOfYear().Format(TimeFormat)
}

// EndOfYear 当年结束时间
func EndOfYear() string {
	return now.EndOfYear().Format(TimeFormat)
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
	return startTime.Format(TimeFormat)
}

// 获取结束时间如: 20230418 到2023-04-18 23:59:59
func GetEndtTime(timeStr string, layout string) string {
	startTime, _ := time.Parse(layout, timeStr)
	endTime := startTime.Add(time.Hour*24 - time.Second)
	return endTime.Format(TimeFormat)
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
	parseTime, err := time.Parse(TimeFormat, currentTime)
	if err != nil {
		return currentTime
	}

	locationTime, err := time.LoadLocation(timeZone)
	if err != nil {
		return currentTime
	}

	return parseTime.In(locationTime).Format(layout)
}

// GetConvertTime 转换时间格式
func GetConvertTime(timeStr, layout, outputLayout string) string {
	// 将时间字符串解析为时间对象
	t, _ := time.Parse(layout, timeStr)
	// 格式化时间为指定格式的字符串
	return t.Format(outputLayout)
}

// 获取月份数据
func GetMonthTime(year int, month time.Month, layout string) (string, string) {
	if month == 0 {
		month = time.Now().Month()
	}
	if year == 0 {
		year = time.Now().Year()
	}
	// Create the start time of the month
	startTime := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	// Create the end time of the month
	endTime := startTime.AddDate(0, 1, -1).Add(time.Hour * 23).Add(time.Minute * 59).Add(time.Second * 59)

	return startTime.Format(layout), endTime.Format(layout)
}

// 解析月份
func ParseMonth(monthStr string) (time.Month, error) {
	monthMap := map[string]time.Month{
		"1":  time.January,
		"01": time.January,
		"2":  time.February,
		"02": time.February,
		"3":  time.March,
		"03": time.March,
		"4":  time.April,
		"04": time.April,
		"5":  time.May,
		"05": time.May,
		"6":  time.June,
		"06": time.June,
		"7":  time.July,
		"07": time.July,
		"8":  time.August,
		"08": time.August,
		"9":  time.September,
		"09": time.September,
		"10": time.October,
		"11": time.November,
		"12": time.December,
	}

	month, ok := monthMap[monthStr]
	if !ok {
		return 0, errors.New(fmt.Sprintf("Invalid month: %s", monthStr))
	}

	return month, nil
}

// 判断是否为当天时间格式
func IsSameDay(timeValue, layout string) bool {
	t, err := time.Parse(layout, timeValue)
	if err != nil {
		return false
	}
	currentDate := time.Now().Format(YMDFormat)
	targetDate := t.Format(YMDFormat)
	return currentDate == targetDate
}

// IsTimestamp 判断是否为时间戳(秒级或毫秒级)
func IsTimestamp(timestamp int64) bool {
	length := len(strconv.FormatInt(timestamp, 10))

	return length == 10 || length == 13
}

// ConvertTimestamp 秒级时间戳转毫秒级时间戳
func ConvertTimestamp(timestamp int64) int64 {
	length := len(strconv.FormatInt(timestamp, 10))
	if length == 10 {
		return timestamp * 1000 // 转换为毫秒级时间戳
	}

	return timestamp // 否则返回原值
}

// ConvertToSeconds 毫秒级时间戳转秒级时间戳
func ConvertToSeconds(timestamp int64) int64 {
	length := len(strconv.FormatInt(timestamp, 10))
	if length == 13 {
		return timestamp / 1000 // 转换为秒级时间戳
	}

	return timestamp // 否则返回原值
}
