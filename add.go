package platform

import (
	"path/filepath"
	"runtime"
)

type Theme struct{}

func (Theme) FolderPath() string {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	return filepath.ToSlash(dir)
}

func (Theme) FunctionMessageName() string {
	return "ShowMessageToUser"
}
