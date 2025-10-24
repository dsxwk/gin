package utils

import (
	"fmt"
	"reflect"
	"strings"
)

// HasKey 检查map键名是否存在(支持任意键类型,但是键类型必须是可比较类型)
func HasKey[K comparable, V any](data map[K]V, key K) bool {
	_, exists := data[key]
	return exists
}

// InArray 检查数组中是否存在某个值
func InArray[T comparable](val T, arr []T) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

// InArrayFast 检查数组中是否存在某个值(高性能适合多次查找场景)
func InArrayFast[T comparable](val T, arr []T) bool {
	m := make(map[T]struct{}, len(arr))
	for _, v := range arr {
		m[v] = struct{}{}
	}
	_, ok := m[val]
	return ok
}

// ArrayUnique 数组去重
func ArrayUnique[T comparable](arr []T) []T {
	m := make(map[T]struct{})
	res := make([]T, 0, len(arr))
	for _, v := range arr {
		if _, exists := m[v]; !exists {
			m[v] = struct{}{}
			res = append(res, v)
		}
	}
	return res
}

// IndexOf 返回元素在切片中的下标,未找到返回-1
func IndexOf[T comparable](val T, arr []T) int {
	for i, v := range arr {
		if v == val {
			return i
		}
	}
	return -1
}

// ArrayFilter 过滤切片,保留满足条件的元素
func ArrayFilter[T any](arr []T, fn func(T) bool) []T {
	res := make([]T, 0)
	for _, v := range arr {
		if fn(v) {
			res = append(res, v)
		}
	}
	return res
}

// StructToMap 将结构体转换为map
func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}

// ArrayToString 将数组格式化为字符串
func ArrayToString(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}
