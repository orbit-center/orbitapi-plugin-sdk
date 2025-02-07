// Package manager 提供插件生命周期管理功能
package manager

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

	"api-plugin-sdk/interfaces"
)

// LifecycleManager 插件生命周期管理器
type LifecycleManager struct {
	plugins     map[string]interfaces.Plugin
	pluginPaths map[string]string // 存储插件文件路径
	mu          sync.RWMutex
}

// NewLifecycleManager 创建生命周期管理器实例
func NewLifecycleManager() *LifecycleManager {
	return &LifecycleManager{
		plugins:     make(map[string]interfaces.Plugin),
		pluginPaths: make(map[string]string),
	}
}

// RegisterPlugin 注册插件
func (m *LifecycleManager) RegisterPlugin(plugin interfaces.Plugin) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	info := plugin.GetInfo()
	if _, exists := m.plugins[info.Name]; exists {
		return fmt.Errorf("插件 %s 已经注册", info.Name)
	}

	m.plugins[info.Name] = plugin
	return nil
}

// IsPluginRegistered 检查插件是否已注册
func (m *LifecycleManager) IsPluginRegistered(name string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	_, exists := m.plugins[name]
	return exists
}

// InstallPlugin 安装插件
func (m *LifecycleManager) InstallPlugin(sourcePath, destPath string) error {
	// 校验插件文件
	if err := m.validatePlugin(sourcePath); err != nil {
		return fmt.Errorf("插件验证失败: %w", err)
	}

	// 复制插件文件到目标目录
	if err := m.copyPlugin(sourcePath, destPath); err != nil {
		return fmt.Errorf("插件复制失败: %w", err)
	}

	return nil
}

// StartPlugin 启动插件
func (m *LifecycleManager) StartPlugin(name string, config interfaces.PluginConfig) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	plugin, exists := m.plugins[name]
	if !exists {
		return fmt.Errorf("插件 %s 未注册", name)
	}

	if err := plugin.Initialize(config); err != nil {
		return fmt.Errorf("初始化插件失败: %w", err)
	}

	if err := plugin.Start(); err != nil {
		return fmt.Errorf("启动插件失败: %w", err)
	}

	return nil
}

// StopPlugin 停止插件
func (m *LifecycleManager) StopPlugin(name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	plugin, exists := m.plugins[name]
	if !exists {
		return fmt.Errorf("插件 %s 未注册", name)
	}

	if err := plugin.Stop(); err != nil {
		return fmt.Errorf("停止插件失败: %w", err)
	}

	return nil
}

// UninstallPlugin 卸载插件
func (m *LifecycleManager) UninstallPlugin(name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.plugins[name]; !exists {
		return fmt.Errorf("插件 %s 未注册", name)
	}

	delete(m.plugins, name)
	delete(m.pluginPaths, name)
	return nil
}

// validatePlugin 验证插件文件
func (m *LifecycleManager) validatePlugin(path string) error {
	// 检查文件是否存在
	if _, err := os.Stat(path); err != nil {
		return fmt.Errorf("插件文件不存在: %w", err)
	}

	// 计算文件哈希
	hash, err := m.calculateFileHash(path)
	if err != nil {
		return fmt.Errorf("计算文件哈希失败: %w", err)
	}

	// TODO: 验证插件签名和兼容性
	_ = hash

	return nil
}

// copyPlugin 复制插件文件到目标目录
func (m *LifecycleManager) copyPlugin(src, dest string) error {
	// 创建目标目录
	if err := os.MkdirAll(filepath.Dir(dest), 0755); err != nil {
		return err
	}

	// 复制文件
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destination.Close()

	if _, err := io.Copy(destination, source); err != nil {
		return err
	}

	return nil
}

// calculateFileHash 计算文件哈希值
func (m *LifecycleManager) calculateFileHash(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
