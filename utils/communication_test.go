package utils

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"api-plugin-sdk/interfaces"
)

func TestHTTPCommunicator(t *testing.T) {
	// 创建测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/message":
			json.NewEncoder(w).Encode(interfaces.Response{
				Success: true,
				Data: map[string]interface{}{
					"message": "ok",
				},
			})
		case "/api/subscribe":
			json.NewEncoder(w).Encode([]interfaces.Message{
				{
					Type:  "test",
					Topic: "test-topic",
					Payload: map[string]interface{}{
						"message": "test message",
					},
				},
			})
		}
	}))
	defer server.Close()

	comm := NewHTTPCommunicator(server.URL, "test-token")

	// 测试发送消息
	t.Run("SendMessage", func(t *testing.T) {
		msg := interfaces.Message{
			Type:  "test",
			Topic: "test-topic",
			Payload: map[string]interface{}{
				"message": "hello",
			},
		}

		resp, err := comm.SendMessage(context.Background(), msg)
		if err != nil {
			t.Errorf("发送消息失败: %v", err)
		}

		if !resp.Success {
			t.Error("响应应该表示成功")
		}
	})

	// 测试订阅
	t.Run("Subscribe", func(t *testing.T) {
		ch, err := comm.Subscribe(context.Background(), "test-topic")
		if err != nil {
			t.Errorf("订阅失败: %v", err)
		}

		select {
		case msg := <-ch:
			if msg.Type != "test" {
				t.Errorf("消息类型不匹配: 期望 test, 实际 %s", msg.Type)
			}
		case <-time.After(2 * time.Second):
			t.Error("接收消息超时")
		}
	})
}
