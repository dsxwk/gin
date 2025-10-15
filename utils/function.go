package utils

import (
	"os"
	"path/filepath"
	"strings"
	"unicode"
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

// UcFirst 首字母大写
func UcFirst(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// LcFirst 首字母小写
func LcFirst(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}
