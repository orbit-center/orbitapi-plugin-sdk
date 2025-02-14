// Package config 提供配置管理功能
package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// ConfigManager 配置管理器
type ConfigManager struct {
	configDir string
}

// NewConfigManager 创建配置管理器实例
func NewConfigManager(configDir string) *ConfigManager {
	return &ConfigManager{
		configDir: configDir,
	}
}

// Config 插件配置结构
type Config struct {
	BaseURL string                 `json:"base_url"` // API基础地址
	Token   string                 `json:"token"`    // 认证令牌
	Params  map[string]interface{} `json:"params"`   // 可选参数配置
	Debug   bool                   `json:"debug"`    // 是否开启调试模式
}

// LoadConfig 加载指定环境的配置
func (m *ConfigManager) LoadConfig(env string) (*Config, error) {
	configFile := filepath.Join(m.configDir, fmt.Sprintf("config.%s.json", env))
	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}

	return &config, nil
}

// SaveConfig 保存配置到指定环境
func (m *ConfigManager) SaveConfig(env string, config *Config) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("序列化配置失败: %w", err)
	}

	configFile := filepath.Join(m.configDir, fmt.Sprintf("config.%s.json", env))
	if err := os.WriteFile(configFile, data, 0644); err != nil {
		return fmt.Errorf("写入配置文件失败: %w", err)
	}

	return nil
}
