// Package client 提供基础的HTTP客户端实现
package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/orbit-center/sdk/interfaces"
	"github.com/orbit-center/sdk/models"
)

// RequestFunc 定义了请求处理函数类型
type RequestFunc func(*http.Request) (*http.Response, error)

// Client 提供与子站点API交互的HTTP客户端
// 负责处理所有HTTP请求，包括认证、序列化和错误处理
type Client struct {
	baseURL     string       // API基础URL
	httpClient  *http.Client // HTTP客户端实例
	token       string       // 认证令牌
	middlewares []Middleware
}

// NewClient 创建新的客户端实例
// baseURL: API基础URL，例如 "http://api.example.com/v1"
// token: 认证令牌
func NewClient(baseURL, token string) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: &http.Client{},
		token:      token,
	}
}

// Use 添加中间件到客户端
func (c *Client) Use(m ...Middleware) {
	c.middlewares = append(c.middlewares, m...)
}

// doRequest 执行HTTP请求
// method: HTTP方法，如 "GET"、"POST" 等
// path: API路径，如 "/users/profile"
// body: 请求体，会被序列化为JSON
// 返回响应体的字节数组和可能的错误
func (c *Client) doRequest(method, path string, body interface{}) ([]byte, error) {
	resp, err := c.DoRequest(method, path, body)
	if err != nil {
		return nil, err
	}

	var data interface{}
	if err := resp.DecodeData(&data); err != nil {
		return nil, err
	}

	return json.Marshal(data)
}

// User 返回用户服务实例
func (c *Client) User() interfaces.UserService {
	return &UserService{Client: c}
}

// Role 返回角色服务实例
func (c *Client) Role() interfaces.RoleService {
	return &RoleService{Client: c}
}

// Menu 返回菜单服务实例
func (c *Client) Menu() interfaces.MenuService {
	return &MenuService{Client: c}
}

// Dict 返回字典服务实例
func (c *Client) Dict() interfaces.DictService {
	return &DictService{Client: c}
}

// DoRequest 执行HTTP请求并返回响应
func (c *Client) DoRequest(method, path string, body interface{}) (interfaces.Response, error) {
	var reqBody []byte
	var err error

	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal request body failed: %w", err)
		}
	}

	req, err := http.NewRequest(method, c.baseURL+path, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.token)

	// 构建请求处理链
	var handler RequestFunc = c.httpClient.Do
	for i := len(c.middlewares) - 1; i >= 0; i-- {
		handler = c.middlewares[i](handler)
	}

	resp, err := handler(req)
	if err != nil {
		return nil, fmt.Errorf("do request failed: %w", err)
	}

	// 创建一个新的响应对象，由它负责关闭响应体
	return &response{resp: resp}, nil
}

// response 实现 interfaces.Response 接口
type response struct {
	resp *http.Response
	body []byte // 缓存已读取的响应体
}

// DecodeData 实现 interfaces.Response 接口
func (r *response) DecodeData(v interface{}) error {
	var err error
	if r.body == nil {
		r.body, err = io.ReadAll(r.resp.Body)
		r.resp.Body.Close()
		if err != nil {
			return fmt.Errorf("read response body failed: %w", err)
		}
	}

	// 先检查 HTTP 状态码
	if r.resp.StatusCode >= 400 {
		apiErr := &APIError{
			Code:    r.resp.StatusCode,
			Message: r.resp.Status,
		}

		// 尝试解析错误响应
		var resp models.Response
		if json.Unmarshal(r.body, &resp) == nil {
			apiErr.Code = resp.Code
			apiErr.Message = resp.Message
		}
		return apiErr
	}

	// 解析标准响应
	var resp models.Response
	if err := json.Unmarshal(r.body, &resp); err != nil {
		return fmt.Errorf("unmarshal response failed: %w", err)
	}

	// 使用 mapstructure 进行类型转换
	config := &mapstructure.DecoderConfig{
		Result:           v,
		TagName:          "json",
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			mapstructure.StringToSliceHookFunc(","),
		),
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return fmt.Errorf("create decoder failed: %w", err)
	}

	return decoder.Decode(resp.Data)
}

// Error 实现 error 接口
func (e *APIError) Error() string {
	return fmt.Sprintf("API error: code=%d, message=%s", e.Code, e.Message)
}
