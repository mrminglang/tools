package times_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/mrminglang/tools/dumps"
	"github.com/mrminglang/tools/times"
	"github.com/stretchr/testify/assert"
)

func TestStartOfDay(t *testing.T) {
	dayStartTime := times.StartOfDay()
	dumps.Dump(dayStartTime)
}

func TestEndOfDay(t *testing.T) {
	dayEndTime := times.EndOfDay()
	dumps.Dump(dayEndTime)
}

func TestStartOfMonth(t *testing.T) {
	monthStartTime := times.StartOfMonth()
	dumps.Dump(monthStartTime)
}

func TestEndOfMonth(t *testing.T) {
	monthEndTime := times.EndOfMonth()
	dumps.Dump(monthEndTime)
}

func TestStartOfYear(t *testing.T) {
	yearStartTime := times.StartOfYear()
	dumps.Dump(yearStartTime)
}

func TestEndOfYear(t *testing.T) {
	yearEndTime := times.EndOfYear()
	dumps.Dump(yearEndTime)
}

func TestGetYearMonth(t *testing.T) {
	dumps.Dump(times.GetYearMonth(times.YYYYMM))
}

func TestGetYearProximo(t *testing.T) {
	dumps.Dump(times.GetYearProximo(times.YYYYMM))
}

func TestGetStartTime(t *testing.T) {
	dumps.Dump(times.GetStartTime("20230418", times.YYYYMMDD))
}

func TestGetEndtTime(t *testing.T) {
	dumps.Dump(times.GetEndTime("20230418", times.YYYYMMDD))
}

func TestGenerateYearMonthIds(t *testing.T) {
	start := "202303"
	layout := times.YYYYMM
	yearMonthIds := times.GenerateYearMonthIds(start, layout)
	fmt.Println(yearMonthIds)
}

func TestGetYesterday(t *testing.T) {
	layout := times.YMDFormat
	yesterday := times.GetYesterday(layout)
	dumps.Dump(yesterday)
}

func TestGetTimeZone(t *testing.T) {
	currentTime := "2023-05-15 09:33:14"
	timeZone := times.TimeZoneSH
	layout := times.TimeFormat
	dumps.Dump(times.GetTimeZone(currentTime, timeZone, layout))
}

func TestGetConvertTime(t *testing.T) {
	timeStr := "220321183549"
	layout := times.YYMMDDhhmmss
	outputLayout := times.YMDhmsFormat3

	dumps.Dump(times.GetConvertTime(timeStr, layout, outputLayout))
}

func TestGetMonthTime(t *testing.T) {
	monthStartTime, monthEndTime := times.GetMonthTime(2022, 8, times.YYMMDD)
	dumps.Dump(monthStartTime)
	dumps.Dump(monthEndTime)
}

func TestParseMonth(t *testing.T) {
	month, err := times.ParseMonth("8")
	if err != nil {
		assert.Error(t, err)
	}
	dumps.Dump(month)
}

func TestIsSameDay(t *testing.T) {
	timeValue := "2023-11-03 13:38:11"
	ok := times.IsSameDay(timeValue, times.TimeFormat)
	dumps.Dump(ok)
}

func TestIsTimestamp(t *testing.T) {
	dumps.Dump(times.IsTimestamp(123456789))
	dumps.Dump(times.IsTimestamp(1633036800))
	dumps.Dump(times.IsTimestamp(1633036800000))
}

func TestConvertTimestamp(t *testing.T) {
	dumps.Dump(times.ConvertTimestamp(1726712577))
	dumps.Dump(times.ConvertTimestamp(1633036800000))
}

func TestConvertToSeconds(t *testing.T) {
	dumps.Dump(times.ConvertToSeconds(1633036800000))
}

func TestGetFormatTimeStr(t *testing.T) {
	dumps.Dump(times.GetFormatTimeStr(time.Now(), ""))
	dumps.Dump(times.GetFormatTimeStr(time.Now(), times.YYYYMMDD))
}
