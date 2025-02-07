package manager

import (
	"fmt"
	"testing"

	"api-plugin-sdk/interfaces"
)

// MockPlugin 用于测试的模拟插件
type MockPlugin struct {
	initialized bool
	running     bool
	stopped     bool
	name        string
	version     string
	status      string
}

func (p *MockPlugin) Initialize(config interfaces.PluginConfig) error {
	p.initialized = true
	return nil
}

func (p *MockPlugin) Start() error {
	if !p.initialized {
		return fmt.Errorf("插件未初始化")
	}
	p.running = true
	p.status = "active"
	return nil
}

func (p *MockPlugin) Stop() error {
	p.running = false
	p.stopped = true
	p.status = "inactive"
	return nil
}

func (p *MockPlugin) GetInfo() interfaces.PluginInfo {
	return interfaces.PluginInfo{
		Name:    p.name,
		Version: p.version,
		Status:  p.status,
	}
}

func TestLifecycleManager(t *testing.T) {
	manager := NewLifecycleManager()

	// 创建一个测试插件
	plugin := &MockPlugin{
		name:    "Mock Plugin",
		version: "1.0.0",
		status:  "inactive",
	}

	// 测试注册插件
	t.Run("RegisterPlugin", func(t *testing.T) {
		err := manager.RegisterPlugin(plugin)
		if err != nil {
			t.Errorf("注册插件失败: %v", err)
		}

		// 验证插件是否已注册
		if !manager.IsPluginRegistered(plugin.GetInfo().Name) {
			t.Error("插件未成功注册")
		}
	})

	// 测试启动插件
	t.Run("StartPlugin", func(t *testing.T) {
		startPlugin := &MockPlugin{
			name:    "Start Plugin",
			version: "1.0.0",
			status:  "inactive",
		}

		err := manager.RegisterPlugin(startPlugin)
		if err != nil {
			t.Fatalf("注册插件失败: %v", err)
		}

		config := interfaces.PluginConfig{
			Name:    startPlugin.name,
			Version: startPlugin.version,
		}

		err = manager.StartPlugin(startPlugin.name, config)
		if err != nil {
			t.Errorf("启动插件失败: %v", err)
		}

		if startPlugin.status != "active" {
			t.Error("插件状态未更新为 active")
		}
	})

	// 测试停止插件
	t.Run("StopPlugin", func(t *testing.T) {
		// 使用新的插件实例
		stopPlugin := &MockPlugin{
			name:    "Stop Plugin",
			version: "1.0.0",
			status:  "active",
		}

		err := manager.RegisterPlugin(stopPlugin)
		if err != nil {
			t.Fatalf("注册插件失败: %v", err)
		}

		err = manager.StopPlugin(stopPlugin.GetInfo().Name)
		if err != nil {
			t.Errorf("停止插件失败: %v", err)
		}

		if stopPlugin.status != "inactive" {
			t.Error("插件状态未更新为 inactive")
		}
	})

	// 测试卸载插件
	t.Run("UninstallPlugin", func(t *testing.T) {
		// 使用新的插件实例
		uninstallPlugin := &MockPlugin{
			name:    "Uninstall Plugin",
			version: "1.0.0",
			status:  "inactive",
		}

		err := manager.RegisterPlugin(uninstallPlugin)
		if err != nil {
			t.Fatalf("注册插件失败: %v", err)
		}

		err = manager.UninstallPlugin(uninstallPlugin.GetInfo().Name)
		if err != nil {
			t.Errorf("卸载插件失败: %v", err)
		}

		if manager.IsPluginRegistered(uninstallPlugin.GetInfo().Name) {
			t.Error("插件未成功卸载")
		}
	})
}
