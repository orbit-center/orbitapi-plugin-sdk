// Package utils 提供通用工具函数
package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Logger 日志记录器
type Logger struct {
	logger *log.Logger
}

// NewLogger 创建日志记录器实例
func NewLogger(logPath string) (*Logger, error) {
	// 确保日志目录存在
	if err := os.MkdirAll(filepath.Dir(logPath), 0755); err != nil {
		return nil, fmt.Errorf("创建日志目录失败: %w", err)
	}

	// 打开日志文件
	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("打开日志文件失败: %w", err)
	}

	return &Logger{
		logger: log.New(f, "", log.LstdFlags),
	}, nil
}
