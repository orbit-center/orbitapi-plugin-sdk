// Package interfaces 定义认证接口
package interfaces

import "net/http"

// AuthInfo 认证信息
type AuthInfo struct {
	Token     string            `json:"token"`
	ExpiresAt int64             `json:"expires_at"`
	Claims    map[string]string `json:"claims"`
}

// Authenticator 认证接口
type Authenticator interface {
	GetAuthInfo() AuthInfo
	ValidateRequest(r *http.Request) error
	RefreshToken() error
}
