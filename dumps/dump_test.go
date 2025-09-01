package dumps_test

import (
	"testing"

	"github.com/mrminglang/tools/dumps"
)

func TestDump(t *testing.T) {
	dumps.Dump("test")
	dumps.Dump([]string{"test2", "test3"})
}
