package paths_test

import (
	"github.com/mrminglang/tools/dumps"
	"github.com/mrminglang/tools/paths"
	"testing"
)

func TestFullPath(t *testing.T) {
	path := paths.FullPath("paths/paths.go")
	dumps.Dump(path)
}
