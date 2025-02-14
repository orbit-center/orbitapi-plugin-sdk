# 测试指南

## 单元测试

运行所有测试：
```bash
go test ./... -v
```

### 表驱动测试示例

```go
func TestGetUserInfo(t *testing.T) {
    tests := []struct {
        name     string
        response interface{}
        wantErr  bool
    }{
        {
            name: "success",
            response: models.Response{
                Code: 0,
                Data: models.User{ID: 1},
            },
            wantErr: false,
        },
        {
            name: "error",
            response: models.Response{
                Code:    500,
                Message: "internal error",
            },
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // ... test implementation
        })
    }
}
```

## 集成测试

运行集成测试：
```bash
go test ./tests/integration -v
```

### Mock 服务器

使用 httptest 包创建 mock 服务器：
```go
ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(models.Response{
        Code: 0,
        Data: models.User{ID: 1},
    })
}))
defer ts.Close()
```

## 性能测试

运行基准测试：
```bash
go test -bench=. ./...
```

### 基准测试示例

```go
func BenchmarkGetUserInfo(b *testing.B) {
    ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        json.NewEncoder(w).Encode(models.Response{
            Code: 0,
            Data: models.User{ID: 1},
        })
    }))
    defer ts.Close()

    client := NewClient(ts.URL, "test-token")
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        _, _ = client.User().GetUserInfo()
    }
}
```

## 代码覆盖率

生成覆盖率报告：
```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
``` 