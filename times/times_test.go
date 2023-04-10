package times_test

import (
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
	dumps.Dump(times.GetYearMonth())
}
