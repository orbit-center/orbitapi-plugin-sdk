package main

import (
	"fmt"
	"log"

	"github.com/orbit-center/sdk/examples/pkg"
)

func main() {
	// 运行中间件示例
	err := pkg.RunMiddlewareExample("http://localhost:8080", "test-token")
	if err != nil {
		log.Fatalf("运行中间件示例失败: %v", err)
	}
	fmt.Println("中间件示例运行成功")
}
