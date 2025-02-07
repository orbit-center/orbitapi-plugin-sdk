# API Plugin SDK

[![Go Report Card](https://goreportcard.com/badge/github.com/orbit-center/orbitapi-plugin-sdk)](https://goreportcard.com/report/github.com/orbit-center/orbitapi-plugin-sdk)
[![GoDoc](https://godoc.org/github.com/orbit-center/orbitapi-plugin-sdk?status.svg)](https://godoc.org/github.com/orbit-center/orbitapi-plugin-sdk)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

API Plugin SDK 是一个用于开发和管理插件的工具包，支持插件的生命周期管理、配置管理、通信等核心功能。

## 特性

- 插件生命周期管理（安装、启动、停止、卸载）
- HTTP 通信协议支持
- 配置管理与热更新
- 安全认证与加密通信
- 版本管理与兼容性检查

## 快速开始

### 安装

```bash
go get github.com/orbit-center/orbitapi-plugin-sdk
```

### 基本使用

```go
package main

import (
    "github.com/orbit-center/orbitapi-plugin-sdk/interfaces"
    "github.com/orbit-center/orbitapi-plugin-sdk/manager"
)

func main() {
    // 创建插件管理器
    manager := manager.NewLifecycleManager()

    // 注册插件
    plugin := NewMyPlugin()
    err := manager.RegisterPlugin(plugin)
    if err != nil {
        log.Fatal(err)
    }

    // 启动插件
    err = manager.StartPlugin(plugin.GetInfo().Name, config)
    if err != nil {
        log.Fatal(err)
    }
}
```

更多示例请查看 [examples](examples/) 目录。

## 文档

- [集成指南](docs/integration_guide.md)
- [API 文档](docs/api.md)
- [配置管理](docs/plugin_config_management.md)
- [通信协议](docs/plugin_communication.md)

## 示例项目

- [基础插件示例](examples/basic-plugin/)
- [HTTP 通信示例](examples/http-plugin/)
- [配置管理示例](examples/config-plugin/)

## 贡献指南

1. Fork 项目
2. 创建特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建 Pull Request

## 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件

## 维护者

- 维护者：Orbit Center Team
- 项目主页：https://github.com/orbit-center/orbitapi-plugin-sdk

## 致谢

感谢所有贡献者对项目的支持！
