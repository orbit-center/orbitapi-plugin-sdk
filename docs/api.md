# API Plugin SDK API 文档

## 核心接口

### Plugin 接口

插件必须实现以下接口：

```go
type Plugin interface {
    Initialize(config PluginConfig) error  // 初始化插件
    Start() error                         // 启动插件
    Stop() error                          // 停止插件
    GetInfo() PluginInfo                  // 获取插件信息
}
```

### 配置接口

```go
type PluginConfig struct {
    Name    string                 `json:"name"`    // 插件名称
    Version string                 `json:"version"` // 插件版本
    Params  map[string]interface{} `json:"params"`  // 插件参数
}

type PluginInfo struct {
    Name    string // 插件名称
    Version string // 插件版本
    Status  string // 插件状态
}
```

### 通信接口

```go
type Communicator interface {
    // 发送消息
    SendMessage(ctx context.Context, message Message) (Response, error)
    // 订阅主题
    Subscribe(ctx context.Context, topic string) (<-chan Message, error)
    // 取消订阅
    Unsubscribe(ctx context.Context, topic string) error
}

type Message struct {
    Type    string                 `json:"type"`    // 消息类型
    Topic   string                 `json:"topic"`   // 消息主题
    Payload map[string]interface{} `json:"payload"` // 消息内容
}

type Response struct {
    Success bool                   `json:"success"` // 是否成功
    Data    map[string]interface{} `json:"data"`    // 响应数据
}
```

## 使用示例

### 1. 创建插件

```go
type MyPlugin struct {
    name    string
    version string
    status  string
}

func (p *MyPlugin) Initialize(config PluginConfig) error {
    p.name = config.Name
    p.version = config.Version
    p.status = "inactive"
    return nil
}

// ... 实现其他接口方法
```

### 2. 注册和管理插件

```go
manager := NewLifecycleManager()
plugin := &MyPlugin{}

// 注册插件
err := manager.RegisterPlugin(plugin)
if err != nil {
    log.Fatal(err)
}

// 启动插件
err = manager.StartPlugin(plugin.GetInfo().Name, config)
if err != nil {
    log.Fatal(err)
}
```

### 3. 使用通信功能

```go
comm := NewHTTPCommunicator(baseURL, authToken)

// 发送消息
resp, err := comm.SendMessage(ctx, Message{
    Type: "event",
    Topic: "user.created",
    Payload: map[string]interface{}{
        "user_id": "123",
    },
})

// 订阅消息
ch, err := comm.Subscribe(ctx, "notifications")
for msg := range ch {
    // 处理消息
}
```

## 错误处理

所有接口方法都应该返回详细的错误信息：

```go
if err != nil {
    return fmt.Errorf("操作失败: %w", err)
}
```

## 最佳实践

1. 始终实现完整的错误处理
2. 使用 context 控制超时
3. 正确管理资源生命周期
4. 实现优雅关闭
5. 提供详细的日志信息

## 版本兼容性

- API 版本遵循语义化版本规范
- 主版本更新可能包含不兼容的 API 更改
- 次版本更新添加向后兼容的功能
- 修订版本包含向后兼容的错误修复 