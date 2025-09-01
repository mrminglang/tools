package compares_test

import (
	"testing"

	"github.com/mrminglang/tools/compares"
	"github.com/mrminglang/tools/dumps"
)

func TestGetCompares(t *testing.T) {
	str := compares.GetCompares("qqqqq $gt aaaa $gte")
	dumps.Dump(str)
}
