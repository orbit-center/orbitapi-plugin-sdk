// Package interfaces 定义了插件 SDK 的核心接口
package interfaces

// Plugin 定义了插件必须实现的标准接口
type Plugin interface {
	// Initialize 初始化插件，加载配置
	Initialize(config PluginConfig) error

	// Start 启动插件
	Start() error

	// Stop 停止插件
	Stop() error

	// GetInfo 获取插件信息
	GetInfo() PluginInfo
}

// PluginInfo 插件基本信息
type PluginInfo struct {
	Name    string            // 插件名称
	Version string            // 插件版本
	Status  string            // 插件状态
	Meta    map[string]string // 插件元数据
}

// PluginConfig 插件配置结构
type PluginConfig struct {
	Name       string                 // 插件名称
	Version    string                 // 插件版本
	Params     map[string]interface{} // 插件参数
	ConfigPath string                 // 配置文件路径
}
