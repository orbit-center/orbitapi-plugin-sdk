# 认证机制

## Token 认证

SDK 使用 Token 进行认证，支持两种配置方式：

### 1. 通过配置文件

```yaml
# config/config.yaml
auth:
  token: your-token
  expire: 3600  # token过期时间(秒)
```

### 2. 通过环境变量

```bash
export ORBIT_API_TOKEN=your-token
```

### 3. 通过代码设置

```go
client := sdk.NewClient("http://api.example.com", "your-token")
```

## 错误处理

当认证失败时，API 会返回 401 错误：

```go
if err != nil {
    if apiErr, ok := err.(*client.APIError); ok {
        if apiErr.IsUnauthorized() {
            // 处理认证失败
            log.Fatal("认证失败，请检查token")
        }
    }
}
``` 