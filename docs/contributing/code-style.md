# 代码规范

## 命名规范

- 包名使用小写
- 接口名使用 "er" 结尾
- 错误类型使用 "Error" 结尾
- 变量使用驼峰命名

## 注释规范

- 所有导出的类型和方法必须有注释
- 注释应该是完整的句子
- 使用 godoc 格式的注释

示例：
```go
// UserService 提供用户相关的操作
type UserService interface {
    // GetUserInfo 获取当前用户信息
    // 返回用户信息和可能的错误
    GetUserInfo() (*User, error)
}
```

## 错误处理

- 使用 fmt.Errorf 和 %w 包装错误
- 错误信息应该有上下文
- 避免使用 panic

## 测试规范

- 文件名使用 _test.go 结尾
- 测试函数使用 Test 前缀
- 使用表驱动测试
- 包含正常和错误情况的测试 