package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/orbit-center/sdk/models"
)

func BenchmarkClient_DoRequest(b *testing.B) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(models.Response{
			Code:    0,
			Message: "success",
			Data:    map[string]interface{}{"key": "value"},
		})
	}))
	defer ts.Close()

	c := NewClient(ts.URL, "test-token")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := c.doRequest("GET", "/test", nil)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestClient_DoRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/not-found":
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(models.Response{
				Code:    404,
				Message: "not found",
			})
		default:
			json.NewEncoder(w).Encode(models.Response{
				Code: 0,
				Data: map[string]interface{}{
					"id":      1,
					"account": "test",
				},
			})
		}
	}))
	defer ts.Close()

	c := NewClient(ts.URL, "test-token")

	// 测试成功请求
	resp, err := c.DoRequest("GET", "/test", nil)
	if err != nil {
		t.Fatalf("DoRequest failed: %v", err)
	}

	var user models.User
	if err := resp.DecodeData(&user); err != nil {
		t.Fatalf("DecodeData failed: %v", err)
	}

	if user.ID != 1 {
		t.Errorf("expected user ID 1, got %d", user.ID)
	}

	// 测试404错误
	_, err = c.DoRequest("GET", "/not-found", nil)
	if err == nil {
		t.Error("expected error for not found")
	}
	if apiErr, ok := err.(*APIError); !ok || apiErr.Code != 404 {
		t.Errorf("expected APIError with code 404, got %v", err)
	}
}

func TestClient_Middleware(t *testing.T) {
	c := NewClient("http://example.com", "test-token")

	// 测试重试中间件
	retryCount := 0
	c.Use(func(next RequestFunc) RequestFunc {
		return func(req *http.Request) (*http.Response, error) {
			retryCount++
			return next(req)
		}
	})

	// 发起请求触发中间件
	_, _ = c.DoRequest("GET", "/test", nil)
	if retryCount != 1 {
		t.Errorf("middleware not executed, count: %d", retryCount)
	}
}

func TestWithRetry(t *testing.T) {
	c := NewClient("http://example.com", "test-token")
	c.Use(WithRetry(3))

	// 模拟临时错误
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

	c.baseURL = ts.URL
	_, err := c.DoRequest("GET", "/test", nil)
	if err != nil {
		t.Errorf("expected success after retries: %v", err)
	}
	if count != 3 {
		t.Errorf("expected 3 retries, got %d", count)
	}
}

func TestClient_WithTimeout(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		json.NewEncoder(w).Encode(models.Response{Code: 0})
	}))
	defer ts.Close()

	c := NewClient(ts.URL, "test-token")
	c.Use(WithTimeout(50 * time.Millisecond))

	_, err := c.DoRequest("GET", "/test", nil)
	if err == nil {
		t.Error("expected timeout error")
	}
}

func TestClient_WithLogging(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(models.Response{Code: 0})
	}))
	defer ts.Close()

	c := NewClient(ts.URL, "test-token")
	c.Use(WithLogging())

	_, err := c.DoRequest("GET", "/test", nil)
	if err != nil {
		t.Fatalf("DoRequest failed: %v", err)
	}
}

func TestClient_InvalidURL(t *testing.T) {
	c := NewClient("invalid-url", "test-token")
	_, err := c.DoRequest("GET", "/test", nil)
	if err == nil {
		t.Error("expected error for invalid URL")
	}
}
