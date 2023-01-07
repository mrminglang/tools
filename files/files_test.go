package files_test

import (
	"github.com/mrminglang/tools/dumps"
	"github.com/mrminglang/tools/files"
	"github.com/mrminglang/tools/paths"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsFileExist(t *testing.T) {
	path := paths.FullPath("files/files.go")
	logrus.Infoln("path::", path)
	exist, err := files.IsFileExist(path)
	if err != nil {
		assert.Error(t, err)
		return
	}
	dumps.Dump(exist)
	assert.True(t, exist)
}

func TestIsFileNotExist(t *testing.T) {
	path := paths.FullPath("files/files.go")
	exist, err := files.IsFileNotExist(path)
	if err != nil {
		assert.Error(t, err)
		return
	}
	dumps.Dump(exist)
	assert.False(t, exist)
}
