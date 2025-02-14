# Orbit Center SDK

Orbit Center SDK 提供了与子站点进行交互的标准接口，帮助插件快速集成子站点的基础功能。

## 核心特性

- **标准化接口**
  - 统一的认证机制
  - 规范的错误处理
  - 类型安全的 API 调用

- **功能完备**
  - 用户管理（认证、权限、信息管理）
  - 角色管理（角色分配、权限管理）
  - 菜单管理（动态菜单、权限控制）
  - 数据字典（系统参数、业务字典）

- **扩展性**
  - 中间件支持（重试、超时、日志）
  - 多环境配置
  - 自定义错误处理

- **可靠性**
  - 完善的单元测试
  - 集成测试覆盖
  - 生产环境验证

## 快速集成

### 1. 安装

```bash
go get github.com/orbit-center/sdk
```

### 2. 配置

创建配置文件 `configs/config.dev.json`:
```json
{
    "base_url": "http://api.example.com",
    "token": "your-token",
    "params": {
        "retry_count": 3,
        "timeout": "30s",
        "log_level": "info"
    },
    "debug": true
}
```

### 3. 代码集成

```go
package main

import (
    "log"
    "time"
    "github.com/orbit-center/sdk/client"
    "github.com/orbit-center/sdk/config"
)

func main() {
    // 1. 加载配置
    cfg, err := config.NewConfigManager("configs").LoadConfig("dev")
    if err != nil {
        log.Fatal(err)
    }

    // 2. 创建客户端
    c := client.NewClient(cfg.BaseURL, cfg.Token)

    // 3. 配置中间件（可选）
    c.Use(
        client.WithRetry(3),
        client.WithTimeout(30*time.Second),
        client.WithLogging(),
    )

    // 4. 使用服务
    // 用户服务
    userSrv := c.User()
    user, err := userSrv.GetUserInfo()

    // 角色服务
    roleSrv := c.Role()
    roles, err := roleSrv.GetRoleList()

    // 菜单服务
    menuSrv := c.Menu()
    menus, err := menuSrv.GetMenuList()

    // 字典服务
    dictSrv := c.Dict()
    dicts, err := dictSrv.GetDictList()
}
```

## API 服务

### 用户服务 (UserService)
- GetUserInfo() - 获取当前用户信息
- GetUserList() - 获取用户列表
- UpdateUser() - 更新用户信息
- CheckPermission() - 检查权限

### 角色服务 (RoleService)
- GetRoleList() - 获取角色列表
- GetRolePermissions() - 获取角色权限
- AssignRole() - 分配角色

### 菜单服务 (MenuService)
- GetMenuList() - 获取菜单列表
- AddMenu() - 添加菜单
- UpdateMenu() - 更新菜单
- DeleteMenu() - 删除菜单
- CheckMenuPermission() - 检查菜单权限

### 字典服务 (DictService)
- GetDictList() - 获取字典列表
- GetDictTypes() - 获取字典类型
- GetDictContents() - 获取字典内容

## 错误处理

```go
if err != nil {
    switch e := err.(type) {
    case *client.APIError:
        // API 错误（业务错误）
        log.Printf("业务错误: 代码=%d, 消息=%s", e.Code, e.Message)
    default:
        // 系统错误（网络错误、配置错误等）
        log.Printf("系统错误: %v", err)
    }
}
```

## 中间件开发

```go
func CustomMiddleware() client.Middleware {
    return func(next client.RequestFunc) client.RequestFunc {
        return func(req *http.Request) (*http.Response, error) {
            // 请求前处理
            start := time.Now()
            
            resp, err := next(req)
            
            // 请求后处理
            log.Printf("请求耗时: %v", time.Since(start))
            return resp, err
        }
    }
}
```

## 环境要求

- Go 1.18+
- API Server v1.0.0+

## 开发工具

```bash
# 运行测试
go test ./... -v

# 代码检查
golangci-lint run

# 生成文档
godoc -http=:6060
```

## 项目结构

```
sdk/
├── client/          # 客户端实现
│   ├── client.go    # 核心客户端
│   ├── user.go      # 用户服务
│   ├── role.go      # 角色服务
│   ├── menu.go      # 菜单服务
│   └── dict.go      # 字典服务
├── interfaces/      # 接口定义
├── models/          # 数据模型
├── config/          # 配置管理
├── docs/           # 文档
├── examples/       # 示例代码
└── tests/          # 测试代码
```

## 常见问题

1. **配置加载失败**
   - 检查配置文件路径
   - 确认配置文件格式正确
   - 验证必要字段已填写

2. **认证失败**
   - 确认 token 正确且未过期
   - 检查 base_url 配置
   - 验证网络连接

3. **类型转换错误**
   - 确认 API 响应格式
   - 检查模型定义
   - 使用正确的数据类型

## 更多资源

- [API 文档](docs/api.md)
- [最佳实践](docs/best-practices.md)
- [示例代码](examples/)
- [更新日志](CHANGELOG.md)

## 许可证

MIT License


