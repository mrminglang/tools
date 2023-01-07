package paths

import (
	"path/filepath"
	"runtime"
)

// FullPath 获取项目根目录的绝对路径
func FullPath(path string) string {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Join(filepath.Dir(b), "../../../../../")
	return filepath.Join(basePath, path)
}

// FullStoragePath 获取项目storage目录绝对路径
func FullStoragePath(storage string, path string) string {
	return filepath.Join(storage, path)
}
