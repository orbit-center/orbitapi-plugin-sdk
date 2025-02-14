package pkg

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/orbit-center/sdk/models"
)

func TestRunBasicExample(t *testing.T) {
	// 创建测试服务器
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(models.Response{
			Code: 0,
			Data: models.User{ID: 1},
		})
	}))
	defer ts.Close()

	// 运行示例代码
	err := RunBasicExample(ts.URL, "test-token")
	if err != nil {
		t.Fatalf("RunBasicExample failed: %v", err)
	}
}

func TestRunMiddlewareExample(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(models.Response{Code: 0})
	}))
	defer ts.Close()

	err := RunMiddlewareExample(ts.URL, "test-token")
	if err != nil {
		t.Fatalf("RunMiddlewareExample failed: %v", err)
	}
}

func TestRunMiddlewareExample_WithError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	err := RunMiddlewareExample(ts.URL, "test-token")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestRunBasicExample_WithError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.Response{
			Code:    401,
			Message: "unauthorized",
		})
	}))
	defer ts.Close()

	err := RunBasicExample(ts.URL, "invalid-token")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestRunBasicExample_WithJWT(t *testing.T) {
	// 创建测试服务器
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 验证JWT token
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(models.Response{
				Code:    401,
				Message: "unauthorized",
			})
			return
		}

		// 验证通过，返回成功响应
		json.NewEncoder(w).Encode(models.Response{
			Code: 0,
			Data: models.User{
				ID:       1,
				Account:  "test",
				Nickname: "测试用户",
			},
		})
	}))
	defer ts.Close()

	// 使用测试token运行示例
	err := RunBasicExample(ts.URL, "test-token")
	if err != nil {
		t.Fatalf("RunBasicExample failed: %v", err)
	}
}

func TestRunMiddlewareExample_WithInvalidToken(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 验证JWT token
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(models.Response{
				Code:    401,
				Message: "unauthorized",
			})
			return
		}

		// 模拟无效token
		if !strings.HasSuffix(authHeader, "test-token") {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(models.Response{
				Code:    401,
				Message: "invalid token",
			})
			return
		}
	}))
	defer ts.Close()

	err := RunMiddlewareExample(ts.URL, "invalid-token")
	if err == nil {
		t.Fatal("expected error for invalid token")
	}
}

func TestRunMiddlewareExample_WithRetry(t *testing.T) {
	count := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		count++
		if count < 3 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		json.NewEncoder(w).Encode(models.Response{Code: 0})
	}))
	defer ts.Close()

	err := RunMiddlewareExample(ts.URL, "test-token")
	if err != nil {
		t.Fatalf("RunMiddlewareExample failed: %v", err)
	}
	if count != 3 {
		t.Errorf("expected 3 retries, got %d", count)
	}
}
