package platform

import (
	"path/filepath"
	"runtime"
)

type Theme struct{}

func (t Theme) PathTemplateIndexHTML() string {
	return t.FolderPath() + "/layout/platform.html"
}

func (Theme) FolderPath() string {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	return filepath.ToSlash(dir)
}
