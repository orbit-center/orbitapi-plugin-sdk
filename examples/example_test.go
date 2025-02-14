package examples

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/orbit-center/sdk/models"
)

func TestExample_Basic(t *testing.T) {
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

func TestExample_WithMiddleware(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(models.Response{Code: 0})
	}))
	defer ts.Close()

	err := RunMiddlewareExample(ts.URL, "test-token")
	if err != nil {
		t.Fatalf("RunMiddlewareExample failed: %v", err)
	}
}

// RunBasicExample 是一个示例函数，用于演示基本功能。
// 参数：
// - url: 服务器的URL。
// - token: 认证令牌。
// 返回值：
// - error: 如果执行失败，返回错误信息。
func RunBasicExample(url string, token string) error {
	// TODO: 实现函数逻辑
	return nil
}

// RunMiddlewareExample 是一个示例函数，用于演示中间件功能。
// 参数：
// - url: 服务器的URL。
// - token: 认证令牌。
// 返回值：
// - error: 如果执行失败，返回错误信息。
func RunMiddlewareExample(url string, token string) error {
	// TODO: 实现函数逻辑
	return nil
}
