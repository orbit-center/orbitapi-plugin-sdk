package client

// APIError 表示API调用过程中的错误
// 包含错误码和错误信息
type APIError struct {
	Code    int    `json:"code"`    // 错误码
	Message string `json:"message"` // 错误信息
}

// IsNotFound 判断是否为404错误
func (e *APIError) IsNotFound() bool {
	return e.Code == 404
}

// IsUnauthorized 判断是否为401错误
func (e *APIError) IsUnauthorized() bool {
	return e.Code == 401
}
