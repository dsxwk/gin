package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var HttpClient = &http.Client{
	Timeout: 5 * time.Second,
	Transport: &http.Transport{
		MaxIdleConns:        100,              // 设置最大空闲连接数
		MaxIdleConnsPerHost: 100,              // 设置每个主机的最大空闲连接数
		IdleConnTimeout:     90 * time.Second, // 设置空闲连接超时时间
	},
}

// HttpRequest 发送http请求
func HttpRequest(method, url string, body []byte, headers map[string]string) (string, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("创建请求失败: %w", err)
	}

	// 设置默认header
	if headers == nil {
		headers = map[string]string{
			"Content-Type": "application/json",
		}
	}
	if !HasKey(headers, "Content-Type") {
		headers["Content-Type"] = "application/json"
	}

	// 设置header
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// 使用全局客户端
	client := HttpClient
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("请求失败: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println("关闭响应失败: ", err)
		}
	}(resp.Body)

	// 读取响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应失败: %w", err)
	}

	return string(respBody), nil
}

// HttpRequestJson 发送json请求并解析响应为结构体
func HttpRequestJson[T any](method, url string, body []byte, headers map[string]string) (*T, error) {
	// 执行请求
	resp, err := HttpRequest(method, url, body, headers)
	if err != nil {
		return nil, err
	}

	// 解析响应
	var result T
	if err = json.Unmarshal([]byte(resp), &result); err != nil {
		return nil, fmt.Errorf("json解析失败: %w", err)
	}

	return &result, nil
}
