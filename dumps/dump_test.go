package dumps_test

import (
	"github.com/mrminglang/tools/dumps"
	"testing"
)


func TestDump(t *testing.T) {
	dumps.Dump("test")
	dumps.Dump([]string{"test2", "test3"})
}
