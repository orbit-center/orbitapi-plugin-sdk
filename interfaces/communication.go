// Package interfaces 定义了插件 SDK 的核心接口
package interfaces

import "context"

// Communication 定义了插件通信接口
type Communication interface {
	// SendMessage 发送消息到子站点
	SendMessage(ctx context.Context, message Message) (Response, error)

	// Subscribe 订阅子站点消息
	Subscribe(ctx context.Context, topic string) (<-chan Message, error)

	// Unsubscribe 取消订阅
	Unsubscribe(ctx context.Context, topic string) error
}

// Message 通信消息结构
type Message struct {
	Type    string                 `json:"type"`
	Topic   string                 `json:"topic"`
	Payload map[string]interface{} `json:"payload"`
}

// Response 响应结构
type Response struct {
	Success bool                   `json:"success"`
	Data    map[string]interface{} `json:"data,omitempty"`
	Error   string                 `json:"error,omitempty"`
}
