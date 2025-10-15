package utils

import (
	"os"
	"path/filepath"
	"strconv"
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

// Integer 整数类型
type Integer interface {
	int | int8 | int16 | int32 | int64
}

// StringToInt 字符串转整数
func StringToInt[T Integer](s string) (T, error) {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		var zero T
		return zero, err
	}
	return T(val), nil
}

// IntToString 整数转字符串
func IntToString[T Integer](v T) string {
	return strconv.FormatInt(int64(v), 10)
}
