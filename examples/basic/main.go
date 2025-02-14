package main

import (
	"fmt"
	"log"

	"github.com/orbit-center/sdk/examples/pkg"
)

func main() {
	// 运行基础示例
	err := pkg.RunBasicExample("http://localhost:8080", "test-token")
	if err != nil {
		log.Fatalf("运行基础示例失败: %v", err)
	}
	fmt.Println("基础示例运行成功")
}
