# 中间件使用

## 内置中间件

SDK 提供了以下内置中间件：

### 重试中间件

```go
client.Use(client.WithRetry(3))  // 最多重试3次
```

### 日志中间件

```go
client.Use(client.WithLogging())
```

## 自定义中间件

可以自定义中间件来扩展功能：

```go
func WithTimeout(timeout time.Duration) client.Middleware {
    return func(req *http.Request) (*http.Response, error) {
        ctx, cancel := context.WithTimeout(req.Context(), timeout)
        defer cancel()
        
        req = req.WithContext(ctx)
        return http.DefaultClient.Do(req)
    }
}

// 使用
client.Use(WithTimeout(5 * time.Second))
```

## 中间件顺序

中间件按照添加顺序逆序执行，例如：

```go
client.Use(
    client.WithLogging(),    // 第三个执行
    client.WithRetry(3),     // 第二个执行
    WithTimeout(5 * time.Second), // 第一个执行
)
``` 