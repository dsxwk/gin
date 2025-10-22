package utils

import (
	"os"
	"path/filepath"
	"strings"
)

// GetRootPath 获取项目根路径
func GetRootPath() string {
	var (
		rootPath string
	)

	execPath, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	rootPath, err = filepath.Abs(execPath)
	if err != nil {
		panic(err.Error())
	}

	return strings.ReplaceAll(rootPath, "\\", "/")
}
