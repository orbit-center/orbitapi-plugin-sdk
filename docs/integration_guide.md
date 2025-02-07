# API Plugin SDK 集成指南

本指南将帮助您将 API Plugin SDK 集成到您的项目中，并开发自定义插件。

## 目录

1. [安装与配置](#安装与配置)
2. [开发插件](#开发插件)
3. [配置管理](#配置管理)
4. [通信实现](#通信实现)
5. [最佳实践](#最佳实践)
6. [常见问题](#常见问题)

## 安装与配置

### 1. 添加依赖

```bash
go get github.com/your-org/api-plugin-sdk
```

### 2. 导入包

```go
import (
    "github.com/your-org/api-plugin-sdk/interfaces"
    "github.com/your-org/api-plugin-sdk/config"
    "github.com/your-org/api-plugin-sdk/manager"
)
```

## 开发插件

### 1. 实现插件接口

```go
package myplugin

import "github.com/your-org/api-plugin-sdk/interfaces"

type MyPlugin struct {
    name    string
    version string
    status  string
    config  interfaces.PluginConfig
}

// Initialize 初始化插件
func (p *MyPlugin) Initialize(config interfaces.PluginConfig) error {
    p.config = config
    p.name = config.Name
    p.version = config.Version
    p.status = "inactive"
    return nil
}

// Start 启动插件
func (p *MyPlugin) Start() error {
    // 实现插件启动逻辑
    p.status = "active"
    return nil
}

// Stop 停止插件
func (p *MyPlugin) Stop() error {
    // 实现插件停止逻辑
    p.status = "inactive"
    return nil
}

// GetInfo 获取插件信息
func (p *MyPlugin) GetInfo() interfaces.PluginInfo {
    return interfaces.PluginInfo{
        Name:    p.name,
        Version: p.version,
        Status:  p.status,
    }
}
```

### 2. 配置管理

```go
// 创建配置文件 config.dev.json
{
    "name": "my-plugin",
    "version": "1.0.0",
    "params": {
        "api_key": "your-api-key",
        "endpoint": "https://api.example.com",
        "timeout": 30
    }
}

// 加载配置
func LoadPluginConfig() (*config.Config, error) {
    manager := config.NewConfigManager("./configs")
    return manager.LoadConfig("dev")
}
```

### 3. 注册和管理插件

```go
func main() {
    // 创建生命周期管理器
    lifecycleManager := manager.NewLifecycleManager()

    // 创建插件实例
    myPlugin := &MyPlugin{}

    // 注册插件
    err := lifecycleManager.RegisterPlugin(myPlugin)
    if err != nil {
        log.Fatalf("注册插件失败: %v", err)
    }

    // 加载配置
    cfg, err := LoadPluginConfig()
    if err != nil {
        log.Fatalf("加载配置失败: %v", err)
    }

    // 启动插件
    err = lifecycleManager.StartPlugin(myPlugin.GetInfo().Name, interfaces.PluginConfig{
        Name:    cfg.Name,
        Version: cfg.Version,
        Params:  cfg.Params,
    })
    if err != nil {
        log.Fatalf("启动插件失败: %v", err)
    }
}
```

### 4. 实现通信

```go
// 创建 HTTP 通信器
func setupCommunication(baseURL, authToken string) (*utils.HTTPCommunicator, error) {
    comm := utils.NewHTTPCommunicator(baseURL, authToken)

    // 发送消息
    message := interfaces.Message{
        Type:  "event",
        Topic: "user.created",
        Payload: map[string]interface{}{
            "user_id": "123",
            "action": "create",
        },
    }

    resp, err := comm.SendMessage(context.Background(), message)
    if err != nil {
        return nil, fmt.Errorf("发送消息失败: %w", err)
    }

    // 订阅消息
    ch, err := comm.Subscribe(context.Background(), "notifications")
    if err != nil {
        return nil, fmt.Errorf("订阅失败: %w", err)
    }

    // 处理接收到的消息
    go func() {
        for msg := range ch {
            handleMessage(msg)
        }
    }()

    return comm, nil
}
```

## 最佳实践

1. **错误处理**
   - 始终检查并处理错误
   - 提供详细的错误信息
   - 实现错误重试机制

```go
func (p *MyPlugin) Start() error {
    maxRetries := 3
    for i := 0; i < maxRetries; i++ {
        err := p.doStart()
        if err == nil {
            return nil
        }
        log.Printf("启动失败，尝试重试 (%d/%d): %v", i+1, maxRetries, err)
        time.Sleep(time.Second * time.Duration(i+1))
    }
    return fmt.Errorf("启动失败，已重试 %d 次", maxRetries)
}
```

2. **资源管理**
   - 使用 defer 确保资源释放
   - 实现优雅关闭
   - 处理并发安全

```go
func (p *MyPlugin) Stop() error {
    defer func() {
        p.status = "inactive"
        // 清理资源
        p.cleanup()
    }()

    // 优雅关闭
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    return p.gracefulShutdown(ctx)
}
```

3. **配置验证**
   - 验证必要的配置项
   - 提供合理的默认值
   - 支持配置热更新

```go
func (p *MyPlugin) validateConfig(config interfaces.PluginConfig) error {
    if config.Name == "" {
        return fmt.Errorf("插件名称不能为空")
    }
    if config.Version == "" {
        return fmt.Errorf("插件版本不能为空")
    }
    return nil
}
```

## 常见问题

1. **Q: 插件启动失败怎么办？**
   A: 检查以下几点：
   - 配置文件是否正确
   - 依赖服务是否可用
   - 查看详细的错误日志

2. **Q: 如何实现插件热更新？**
   A: 可以通过以下步骤：
   - 监听配置文件变化
   - 实现重载配置的方法
   - 确保状态正确迁移

3. **Q: 如何处理并发请求？**
   A: 使用以下策略：
   - 使用互斥锁保护共享资源
   - 实现请求队列
   - 设置合适的超时时间

## 示例项目

完整的示例项目可以在 [examples](../examples) 目录中找到。

## 支持与帮助

如果您在集成过程中遇到任何问题，请：
1. 查看详细的 [API 文档](./api.md)
2. 提交 Issue 到我们的 [GitHub 仓库](https://github.com/your-org/api-plugin-sdk/issues)
3. 联系技术支持团队 