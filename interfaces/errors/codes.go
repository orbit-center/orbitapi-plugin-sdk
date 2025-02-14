// Package errors 定义错误码常量
package errors

// 插件生命周期相关错误码
const (
	// 初始化相关
	ErrInitFailed    = "PLUGIN_001" // 插件初始化失败
	ErrConfigInvalid = "PLUGIN_002" // 配置验证失败

	// 状态相关
	ErrStateTransition = "STATE_001" // 状态转换失败
	ErrStateInvalid    = "STATE_002" // 非法状态

	// 认证相关
	ErrAuthFailed   = "AUTH_001" // 认证失败
	ErrTokenExpired = "AUTH_002" // Token过期

	// 系统相关
	ErrSystemInternal = "SYS_001" // 系统内部错误
)
