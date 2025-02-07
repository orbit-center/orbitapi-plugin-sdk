package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPClient(t *testing.T) {
	// 创建测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 验证请求头
		if r.Header.Get("Content-Type") != "application/json" {
			t.Error("Content-Type 应该为 application/json")
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader != "Bearer test-token" {
			t.Error("Authorization header 不正确")
		}

		// 返回测试响应
		response := map[string]interface{}{
			"success": true,
			"data": map[string]string{
				"message": "ok",
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// 创建 HTTP 客户端
	client := NewHTTPClient(server.URL, "test-token")

	// 测试 POST 请求
	t.Run("TestPost", func(t *testing.T) {
		requestBody := map[string]string{
			"key": "value",
		}

		resp, err := client.Post("/test", requestBody)
		if err != nil {
			t.Errorf("POST 请求失败: %v", err)
		}

		// 验证响应
		var response map[string]interface{}
		if err := json.Unmarshal(resp, &response); err != nil {
			t.Errorf("解析响应失败: %v", err)
		}

		if !response["success"].(bool) {
			t.Error("响应应该表示成功")
		}
	})
}
