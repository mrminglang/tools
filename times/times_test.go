package times_test

import (
	"fmt"
	"github.com/mrminglang/tools/dumps"
	"github.com/mrminglang/tools/times"
	"testing"
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
	dumps.Dump(times.GetEndtTime("20230418", times.YYYYMMDD))
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
	monthStartTime, monthEndTime := times.GetMonthTime(8, times.YYMMDD)
	dumps.Dump(monthStartTime)
	dumps.Dump(monthEndTime)
}
