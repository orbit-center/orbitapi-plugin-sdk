// Package interfaces 定义了核心接口
package interfaces

// Client 定义了HTTP客户端的接口
type Client interface {
	// DoRequest 执行HTTP请求
	DoRequest(method, path string, body interface{}) (Response, error)
}

// Response 定义了API响应的接口
type Response interface {
	// DecodeData 解析响应数据到指定结构
	DecodeData(v interface{}) error
}
