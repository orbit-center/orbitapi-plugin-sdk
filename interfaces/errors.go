// Package interfaces 定义错误处理接口
package interfaces

// PluginError 插件错误类型
type PluginError struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details,omitempty"`
}

// Error 实现 error 接口
func (e *PluginError) Error() string {
	return e.Message
}

// NewPluginError 创建新的插件错误
func NewPluginError(code string, message string) *PluginError {
	return &PluginError{
		Code:    code,
		Message: message,
		Details: make(map[string]interface{}),
	}
}
