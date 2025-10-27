package utils

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"net/url"
	"strings"
	"time"
)

const timeout = 5 * time.Second

type HttpOption struct {
	Headers map[string]string // 请求头
	Query   map[string]string // get参数
	Form    map[string]string // 表单参数
	Body    []byte            // body参数
	Timeout time.Duration     // 超时时间
}

// buildURL 拼接get请求query参数
func buildUrl(baseURL string, query map[string]string) string {
	if len(query) == 0 {
		return baseURL
	}
	q := url.Values{}
	for k, v := range query {
		q.Set(k, v)
	}
	if strings.Contains(baseURL, "?") {
		return baseURL + "&" + q.Encode()
	}
	return baseURL + "?" + q.Encode()
}

// HttpRequest 发送http请求
func HttpRequest(method, uri string, opt *HttpOption) (string, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	if opt == nil {
		opt = &HttpOption{}
	}
	if opt.Timeout == 0 {
		opt.Timeout = timeout
	}

	method = strings.ToUpper(method)
	uri = buildUrl(uri, opt.Query)
	req.Header.SetMethod(method)
	req.SetRequestURI(uri)
	contentType := "application/json"
	if len(opt.Form) > 0 {
		// x-www-form-urlencoded
		values := url.Values{}
		for k, v := range opt.Form {
			values.Set(k, v)
		}
		req.SetBodyString(values.Encode())
		contentType = "application/x-www-form-urlencoded"
	} else if len(opt.Body) > 0 {
		// body
		req.SetBody(opt.Body)
		contentType = "application/json"
	}

	// header设置
	req.Header.Set("Content-Type", contentType)
	for k, v := range opt.Headers {
		req.Header.Set(k, v)
	}

	// 创建客户端并设置超时
	client := &fasthttp.Client{
		MaxConnsPerHost: 100,     // 最大连接数
		ReadTimeout:     timeout, // 读取超时
		WriteTimeout:    timeout, // 写入超时
	}

	// 发送请求
	if err := client.Do(req, resp); err != nil {
		return "", fmt.Errorf("请求失败: %w", err)
	}

	// 获取响应内容
	respBody := string(resp.Body())

	if resp.StatusCode() != 200 {
		return "", fmt.Errorf("请求失败, 状态码: %d, 响应: %s", resp.Header.StatusCode(), respBody)
	}

	return respBody, nil
}

// HttpRequestJson 发送请求并解析json响应
func HttpRequestJson[T any](method, url string, opt *HttpOption) (*T, error) {
	resp, err := HttpRequest(method, url, opt)
	if err != nil {
		return nil, err
	}

	var result T
	if err = json.Unmarshal([]byte(resp), &result); err != nil {
		return nil, fmt.Errorf("json解析失败: %w\n响应内容:\n%s", err, resp)
	}

	return &result, nil
}
