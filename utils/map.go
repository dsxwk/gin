package utils

// HasKey 检查map键名是否存在
func HasKey(data map[string]interface{}, key string) bool {
	_, exists := data[key]

	return exists
}
