package pkg

import (
	"fmt"
	"time"

	jwtmicro "github.com/orbit-center/jwt-go-micro/middleware"
	"github.com/orbit-center/sdk/client"
)

// TokenConfig JWT配置
type TokenConfig struct {
	SigningMethod string
	SigningKey    []byte
	Expiration    time.Duration
}

// DefaultTokenConfig 默认JWT配置
var DefaultTokenConfig = TokenConfig{
	SigningMethod: "HS256",
	SigningKey:    []byte("your-secret-key"),
	Expiration:    24 * time.Hour,
}

// RunBasicExample 运行基础示例
func RunBasicExample(baseURL, token string) error {
	// 创建客户端
	c := client.NewClient(baseURL, token)

	// 添加JWT认证中间件
	jwtMiddleware := jwtmicro.NewJWTAuthWrapper(
		jwtmicro.WithSigningMethod(DefaultTokenConfig.SigningMethod),
		jwtmicro.WithSigningKey(DefaultTokenConfig.SigningKey),
		jwtmicro.WithTokenExpiration(DefaultTokenConfig.Expiration),
	)

	c.Use(client.WithJWTAuth(jwtMiddleware))

	// 获取用户信息
	user, err := c.User().GetUserInfo()
	if err != nil {
		return fmt.Errorf("get user info failed: %w", err)
	}

	fmt.Printf("当前用户: %+v\n", user)
	return nil
}

// RunMiddlewareExample 运行中间件示例
func RunMiddlewareExample(baseURL, token string) error {
	// 创建客户端
	c := client.NewClient(baseURL, token)

	// 添加JWT认证中间件
	jwtMiddleware := jwtmicro.NewJWTAuthWrapper(
		jwtmicro.WithSigningMethod(DefaultTokenConfig.SigningMethod),
		jwtmicro.WithSigningKey(DefaultTokenConfig.SigningKey),
		jwtmicro.WithTokenExpiration(DefaultTokenConfig.Expiration),
	)

	// 添加中间件
	c.Use(
		client.WithJWTAuth(jwtMiddleware),
		client.WithRetry(3),
		client.WithLogging(),
	)

	// 获取角色列表
	roles, err := c.Role().GetRoleList()
	if err != nil {
		return fmt.Errorf("get role list failed: %w", err)
	}

	fmt.Printf("角色列表: %+v\n", roles)
	return nil
}
