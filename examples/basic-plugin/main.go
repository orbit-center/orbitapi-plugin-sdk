package main

import (
	"log"

	"github.com/orbit-center/orbitapi-plugin-sdk/interfaces"
	"github.com/orbit-center/orbitapi-plugin-sdk/manager"
)

// BasicPlugin 基础插件示例
type BasicPlugin struct {
	name    string
	version string
	status  string
}

func (p *BasicPlugin) Initialize(config interfaces.PluginConfig) error {
	p.name = config.Name
	p.version = config.Version
	p.status = "inactive"
	return nil
}

func (p *BasicPlugin) Start() error {
	p.status = "active"
	return nil
}

func (p *BasicPlugin) Stop() error {
	p.status = "inactive"
	return nil
}

func (p *BasicPlugin) GetInfo() interfaces.PluginInfo {
	return interfaces.PluginInfo{
		Name:    p.name,
		Version: p.version,
		Status:  p.status,
	}
}

func main() {
	// 创建插件管理器
	manager := manager.NewLifecycleManager()

	// 创建插件实例
	plugin := &BasicPlugin{}

	// 注册插件
	err := manager.RegisterPlugin(plugin)
	if err != nil {
		log.Fatalf("注册插件失败: %v", err)
	}

	// 启动插件
	config := interfaces.PluginConfig{
		Name:    "basic-plugin",
		Version: "1.0.0",
		Params: map[string]interface{}{
			"key": "value",
		},
	}

	err = manager.StartPlugin(plugin.GetInfo().Name, config)
	if err != nil {
		log.Fatalf("启动插件失败: %v", err)
	}

	log.Printf("插件已启动: %s", plugin.GetInfo().Status)
}
