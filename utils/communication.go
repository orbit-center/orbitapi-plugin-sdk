// Package utils 提供通用工具函数
package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/orbit-center/orbitapi-plugin-sdk/interfaces"
)

// HTTPCommunicator HTTP 通信实现
type HTTPCommunicator struct {
	client     *HTTPClient
	subscribed map[string]chan interfaces.Message
	mu         sync.RWMutex
}

// NewHTTPCommunicator 创建 HTTP 通信实例
func NewHTTPCommunicator(baseURL, authToken string) *HTTPCommunicator {
	return &HTTPCommunicator{
		client:     NewHTTPClient(baseURL, authToken),
		subscribed: make(map[string]chan interfaces.Message),
	}
}

// SendMessage 发送消息到子站点
func (c *HTTPCommunicator) SendMessage(ctx context.Context, message interfaces.Message) (interfaces.Response, error) {
	resp, err := c.client.Post("/api/message", message)
	if err != nil {
		return interfaces.Response{}, fmt.Errorf("发送消息失败: %w", err)
	}

	var response interfaces.Response
	if err := json.Unmarshal(resp, &response); err != nil {
		return interfaces.Response{}, fmt.Errorf("解析响应失败: %w", err)
	}

	return response, nil
}

// Subscribe 订阅子站点消息
func (c *HTTPCommunicator) Subscribe(ctx context.Context, topic string) (<-chan interfaces.Message, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if ch, exists := c.subscribed[topic]; exists {
		return ch, nil
	}

	ch := make(chan interfaces.Message, 100)
	c.subscribed[topic] = ch

	// 启动长轮询
	go c.pollMessages(ctx, topic, ch)

	return ch, nil
}

// Unsubscribe 取消订阅
func (c *HTTPCommunicator) Unsubscribe(ctx context.Context, topic string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if ch, exists := c.subscribed[topic]; exists {
		close(ch)
		delete(c.subscribed, topic)
	}

	return nil
}

// pollMessages 长轮询获取消息
func (c *HTTPCommunicator) pollMessages(ctx context.Context, topic string, ch chan interfaces.Message) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			resp, err := c.client.Post("/api/subscribe", map[string]string{"topic": topic})
			if err != nil {
				log.Printf("轮询消息失败 [%s]: %v", topic, err)
				continue
			}

			var messages []interfaces.Message
			if err := json.Unmarshal(resp, &messages); err != nil {
				log.Printf("解析消息失败 [%s]: %v", topic, err)
				continue
			}

			for _, msg := range messages {
				select {
				case ch <- msg:
				default:
					log.Printf("警告: 消息通道已满，丢弃消息 [%s]", topic)
				}
			}
		}
	}
}
