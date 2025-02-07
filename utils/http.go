// Package utils 提供通用工具函数
package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// HTTPClient HTTP 客户端封装
type HTTPClient struct {
	baseURL    string
	authToken  string
	httpClient *http.Client
}

// NewHTTPClient 创建 HTTP 客户端实例
func NewHTTPClient(baseURL, authToken string) *HTTPClient {
	return &HTTPClient{
		baseURL:    baseURL,
		authToken:  authToken,
		httpClient: &http.Client{},
	}
}

// Post 发送 POST 请求
func (c *HTTPClient) Post(path string, body interface{}) ([]byte, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("序列化请求体失败: %w", err)
	}

	req, err := http.NewRequest("POST", c.baseURL+path, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	c.setHeaders(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("请求失败，状态码: %d", resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}

	return respBody, nil
}

// setHeaders 设置请求头
func (c *HTTPClient) setHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	if c.authToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.authToken)
	}
}
