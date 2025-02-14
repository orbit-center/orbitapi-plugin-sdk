# 错误处理

## 错误类型

SDK 定义了以下错误类型：

```go
type APIError struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}
```

## 常见错误码

- 401: 未认证或token无效
- 403: 无权限访问
- 404: 资源不存在
- 500: 服务器内部错误

## 错误处理示例

```go
user, err := client.User().GetUserInfo()
if err != nil {
    switch {
    case errors.Is(err, client.ErrUnauthorized):
        log.Fatal("认证失败")
    case errors.Is(err, client.ErrPermissionDenied):
        log.Fatal("无权限访问")
    default:
        log.Printf("获取用户信息失败: %v", err)
    }
}
```

## 自定义错误处理

可以通过中间件自定义错误处理逻辑：

```go
client.Use(func(req *http.Request) (*http.Response, error) {
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        // 自定义错误处理
        return nil, fmt.Errorf("请求失败: %w", err)
    }
    return resp, nil
})
``` 