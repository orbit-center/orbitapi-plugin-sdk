# 安装指南

## 系统要求

- Go 1.21 或更高版本
- golangci-lint (可选，用于代码检查)

## 安装步骤

1. 添加依赖
```bash
go get github.com/orbit-center/sdk
```

2. 安装开发工具
```bash
# Windows
scoop install golangci-lint

# macOS
brew install golangci-lint

# Linux
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
```

3. 验证安装
```bash
go run examples/main.go
```

## 配置说明

1. 创建配置文件
```yaml
# config/config.yaml
server:
  host: api.example.com
  port: 443
```

2. 设置环境变量
```bash
export ORBIT_API_TOKEN=your-token
export ORBIT_API_HOST=api.example.com
``` 