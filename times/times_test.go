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
	dumps.Dump(times.GetYearMonth("200601"))
}

func TestGetYearProximo(t *testing.T) {
	dumps.Dump(times.GetYearProximo("200601"))
}

func TestGetStartTime(t *testing.T) {
	dumps.Dump(times.GetStartTime("20230418", "20060102"))
}

func TestGetEndtTime(t *testing.T) {
	dumps.Dump(times.GetEndtTime("20230418", "20060102"))
}

func TestGenerateYearMonthIds(t *testing.T) {
	start := "202303"
	layout := "200601"
	yearMonthIds := times.GenerateYearMonthIds(start, layout)
	fmt.Println(yearMonthIds)
}

func TestGetYesterday(t *testing.T) {
	layout := "2006-01-02"
	yesterday := times.GetYesterday(layout)
	dumps.Dump(yesterday)
}
