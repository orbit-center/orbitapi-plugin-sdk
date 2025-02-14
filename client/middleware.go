package client

import (
	"context"
	"fmt"
	"net/http"
	"time"

	jwtmicro "github.com/orbit-center/jwt-go-micro/middleware"
)

// Middleware 定义了中间件函数类型
type Middleware func(RequestFunc) RequestFunc

// WithTimeout 返回一个超时中间件
func WithTimeout(timeout time.Duration) Middleware {
	return func(next RequestFunc) RequestFunc {
		return func(req *http.Request) (*http.Response, error) {
			ctx, cancel := context.WithTimeout(req.Context(), timeout)
			defer cancel()

			req = req.WithContext(ctx)
			return next(req)
		}
	}
}

// WithRetry 返回一个重试中间件
func WithRetry(maxRetries int) Middleware {
	return func(next RequestFunc) RequestFunc {
		return func(req *http.Request) (*http.Response, error) {
			var lastErr error
			for i := 0; i <= maxRetries; i++ {
				resp, err := next(req)
				if err == nil && resp.StatusCode < 500 {
					return resp, nil
				}
				if resp != nil {
					resp.Body.Close()
				}
				if err != nil {
					lastErr = err
				} else {
					lastErr = fmt.Errorf("request failed with status: %s", resp.Status)
				}
				if i < maxRetries {
					time.Sleep(time.Duration(i+1) * time.Second)
				}
			}
			return nil, fmt.Errorf("max retries reached: %w", lastErr)
		}
	}
}

// WithLogging 返回一个日志中间件
func WithLogging() Middleware {
	return func(next RequestFunc) RequestFunc {
		return func(req *http.Request) (*http.Response, error) {
			// 记录请求开始时间
			start := time.Now()

			// 执行请求
			resp, err := next(req)

			// 记录请求结束时间和结果
			duration := time.Since(start)
			if err != nil {
				// TODO: 使用正式的日志库
				println("Request failed:", req.URL.Path, "error:", err.Error(), "duration:", duration.String())
			} else {
				println("Request succeeded:", req.URL.Path, "status:", resp.Status, "duration:", duration.String())
			}

			return resp, err
		}
	}
}

// WithJWTAuth 返回一个JWT认证中间件
func WithJWTAuth(jwt *jwtmicro.JWTAuthWrapper) Middleware {
	return func(next RequestFunc) RequestFunc {
		return func(req *http.Request) (*http.Response, error) {
			// 从请求头中获取token
			token := req.Header.Get("Authorization")
			if token != "" && len(token) > 7 {
				token = token[7:] // 去掉 "Bearer " 前缀
			}
			if token != "" {
				req.Header.Set("Authorization", "Bearer "+token)
			}
			return next(req)
		}
	}
}
