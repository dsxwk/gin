package utils

import (
	"path/filepath"
	"runtime"
	"strings"
	"unicode"
)

// GetRootPath 获取项目根路径
func GetRootPath() string {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	rootPath := filepath.Join(basePath, "..")
	rootPath, _ = filepath.Abs(rootPath)

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
