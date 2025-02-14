# 用户服务 API

## 获取用户信息

```go
user, err := client.User().GetUserInfo()
```

### 参数
无

### 返回值
- `*models.User`: 用户信息
- `error`: 错误信息

### 示例
```go
client := sdk.NewClient("http://api.example.com", "your-token")
user, err := client.User().GetUserInfo()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("用户: %+v\n", user)
```

## 检查权限

```go
hasPermission := client.User().CheckPermission("system:user:view")
```

### 参数
- `permission string`: 权限标识符

### 返回值
- `bool`: 是否有权限 