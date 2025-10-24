package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

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

// Spaces 返回指定数量的空格
func Spaces(n int) string {
	return fmt.Sprintf("%*s", n, "")
}

// RandString 生成指定长度的随机字符串
func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
