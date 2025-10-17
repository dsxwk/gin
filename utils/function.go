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

// SnakeToCamel 将下划线命名转换为驼峰命名
func SnakeToCamel(s string) string {
	s = strings.ToLower(s)
	var b strings.Builder
	upperNext := true
	for _, r := range s {
		if r == '_' {
			upperNext = true
			continue
		}
		if upperNext {
			b.WriteRune(unicode.ToUpper(r))
			upperNext = false
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// SnakeToLowerCamel 将下划线命名转换为小驼峰命名
func SnakeToLowerCamel(s string) string {
	if s == "" {
		return ""
	}
	camel := SnakeToCamel(s)
	return strings.ToLower(camel[:1]) + camel[1:]
}

// CamelToSnake 将驼峰命名转换为下划线命名
func CamelToSnake(s string) string {
	var b strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				b.WriteByte('_')
			}
			b.WriteRune(unicode.ToLower(r))
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}
