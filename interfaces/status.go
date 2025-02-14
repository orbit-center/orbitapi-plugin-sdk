// Package interfaces 定义插件状态管理接口
package interfaces

import "time"

// PluginStatus 插件状态类型
type PluginStatus string

const (
	StatusInitializing PluginStatus = "initializing"
	StatusRunning      PluginStatus = "running"
	StatusStopped      PluginStatus = "stopped"
	StatusError        PluginStatus = "error"
)

// StatusRecord 状态记录
type StatusRecord struct {
	Status    PluginStatus `json:"status"`
	Timestamp time.Time    `json:"timestamp"`
	Message   string       `json:"message"`
}

// StatusManager 状态管理接口
type StatusManager interface {
	TransitionTo(status PluginStatus) error
	GetCurrentStatus() PluginStatus
	GetStatusHistory() []StatusRecord
}
