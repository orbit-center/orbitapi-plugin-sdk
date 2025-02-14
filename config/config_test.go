package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

type TestConfig struct {
	Name    string            `json:"name"`
	Version string            `json:"version"`
	Params  map[string]string `json:"params"`
}

func TestConfigManager(t *testing.T) {
	// 创建临时测试目录
	tempDir := t.TempDir()

	// 创建测试配置文件
	testConfig := &Config{
		BaseURL: "http://127.0.0.1:8081/api/v1",
		Token:   "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzk2MDQ3MjMsInVzZXJfaWQiOjF9.SMIVWvf-5AhBgcakRuFrcrqpdfP_l8yUvZbx2wRlBLc",
		Params:  map[string]interface{}{"key": "value"},
	}

	configFile := filepath.Join(tempDir, "config.dev.json")
	configData, err := json.MarshalIndent(testConfig, "", "  ")
	if err != nil {
		t.Fatalf("序列化配置失败: %v", err)
	}

	if err := os.WriteFile(configFile, configData, 0644); err != nil {
		t.Fatalf("写入配置文件失败: %v", err)
	}

	manager := NewConfigManager(tempDir)

	// 测试加载配置
	t.Run("LoadConfig", func(t *testing.T) {
		config, err := manager.LoadConfig("dev")
		if err != nil {
			t.Fatalf("加载配置失败: %v", err)
		}

		if config.BaseURL != testConfig.BaseURL {
			t.Errorf("配置名称不匹配: 期望 %s, 实际 %s", testConfig.BaseURL, config.BaseURL)
		}

		if config.Token != testConfig.Token {
			t.Errorf("配置版本不匹配: 期望 %s, 实际 %s", testConfig.Token, config.Token)
		}

		value, ok := config.Params["key"].(string)
		if !ok || value != "value" {
			t.Errorf("配置参数不匹配: 期望 %s, 实际 %s", "value", value)
		}
	})

	// 测试保存配置
	t.Run("SaveConfig", func(t *testing.T) {
		newConfig := &Config{
			BaseURL: "http://127.0.0.1:8081/api/v1",
			Token:   "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzk2MDQ3MjMsInVzZXJfaWQiOjF9.SMIVWvf-5AhBgcakRuFrcrqpdfP_l8yUvZbx2wRlBLc",
		}

		if err := manager.SaveConfig("prod", newConfig); err != nil {
			t.Fatalf("保存配置失败: %v", err)
		}

		// 验证配置是否正确保存
		savedConfig, err := manager.LoadConfig("prod")
		if err != nil {
			t.Fatalf("加载保存的配置失败: %v", err)
		}

		if !reflect.DeepEqual(savedConfig, newConfig) {
			t.Error("保存的配置与原始配置不匹配")
		}
	})
}
