package paths_test

import (
	"github.com/mrminglang/tools/dumps"
	"github.com/mrminglang/tools/paths"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestFullPath(t *testing.T) {
	path := paths.FullPath("paths/paths.go")
	logrus.Infoln("path::", path)
	dumps.Dump(path)
}
