package compares_test

import (
	"github.com/mrminglang/tools/compares"
	"github.com/mrminglang/tools/dumps"
	"testing"
)

func TestGetCompares(t *testing.T) {
	str := compares.GetCompares("qqqqq $gt aaaa $gte")
	dumps.Dump(str)
}
