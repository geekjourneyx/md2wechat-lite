package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad_Default(t *testing.T) {
	// 创建临时目录
	tmpDir := t.TempDir()

	// 设置临时主目录
	oldHome := os.Getenv("HOME")
	oldConfigDir := configDir
	oldConfigPath := configPath
	defer func() {
		os.Setenv("HOME", oldHome)
		configDir = oldConfigDir
		configPath = oldConfigPath
	}()
	os.Setenv("HOME", tmpDir)
	// 重新计算配置路径
	configDir = filepath.Join(tmpDir, ConfigDir)
	configPath = filepath.Join(configDir, ConfigFile)

	// 测试默认配置（无配置文件）
	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() failed: %v", err)
	}

	if cfg.APIBaseURL != DefaultAPIBaseURL {
		t.Errorf("APIBaseURL = %s, want %s", cfg.APIBaseURL, DefaultAPIBaseURL)
	}
}

func TestSave_Load(t *testing.T) {
	// 创建临时目录
	tmpDir := t.TempDir()

	// 设置临时主目录
	oldHome := os.Getenv("HOME")
	oldConfigDir := configDir
	oldConfigPath := configPath
	defer func() {
		os.Setenv("HOME", oldHome)
		configDir = oldConfigDir
		configPath = oldConfigPath
	}()
	os.Setenv("HOME", tmpDir)
	// 重新计算配置路径
	configDir = filepath.Join(tmpDir, ConfigDir)
	configPath = filepath.Join(configDir, ConfigFile)

	// 保存配置
	cfg := &Config{
		WechatAppID:     "wx_test_appid_example",
		WechatAppSecret: "test_secret",
		APIKey:          "test_api_key",
		APIBaseURL:      "https://test.example.com",
	}

	if err := Save(cfg); err != nil {
		t.Fatalf("Save() failed: %v", err)
	}

	// 检查文件是否存在
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Fatalf("Config file not created: %s", configPath)
	}

	// 加载配置
	loadedCfg, err := Load()
	if err != nil {
		t.Fatalf("Load() failed: %v", err)
	}

	if loadedCfg.WechatAppID != cfg.WechatAppID {
		t.Errorf("WechatAppID = %s, want %s", loadedCfg.WechatAppID, cfg.WechatAppID)
	}
	if loadedCfg.WechatAppSecret != cfg.WechatAppSecret {
		t.Errorf("WechatAppSecret = %s, want %s", loadedCfg.WechatAppSecret, cfg.WechatAppSecret)
	}
	if loadedCfg.APIKey != cfg.APIKey {
		t.Errorf("APIKey = %s, want %s", loadedCfg.APIKey, cfg.APIKey)
	}
}

func TestSet_Get(t *testing.T) {
	// 创建临时目录
	tmpDir := t.TempDir()

	// 设置临时主目录
	oldHome := os.Getenv("HOME")
	oldConfigDir := configDir
	oldConfigPath := configPath
	defer func() {
		os.Setenv("HOME", oldHome)
		configDir = oldConfigDir
		configPath = oldConfigPath
	}()
	os.Setenv("HOME", tmpDir)
	configDir = filepath.Join(tmpDir, ConfigDir)
	configPath = filepath.Join(configDir, ConfigFile)

	// 设置配置
	if err := Set("wechat-appid", "wx_test_appid_example"); err != nil {
		t.Fatalf("Set() failed: %v", err)
	}

	// 获取配置
	value, err := Get("wechat-appid")
	if err != nil {
		t.Fatalf("Get() failed: %v", err)
	}

	if value != "wx_test_appid_example" {
		t.Errorf("value = %s, want wx_test_appid_example", value)
	}
}

func TestEnvironmentOverride(t *testing.T) {
	// 创建临时目录
	tmpDir := t.TempDir()

	// 设置临时主目录
	oldHome := os.Getenv("HOME")
	oldConfigDir := configDir
	oldConfigPath := configPath
	defer func() {
		os.Setenv("HOME", oldHome)
		configDir = oldConfigDir
		configPath = oldConfigPath
	}()
	os.Setenv("HOME", tmpDir)
	configDir = filepath.Join(tmpDir, ConfigDir)
	configPath = filepath.Join(configDir, ConfigFile)

	// 设置配置文件
	cfg := &Config{
		WechatAppID:     "wx_file_value",
		WechatAppSecret: "file_secret",
		APIKey:          "file_key",
		APIBaseURL:      "https://file.example.com",
	}
	if err := Save(cfg); err != nil {
		t.Fatalf("Save() failed: %v", err)
	}

	// 设置环境变量
	os.Setenv("MD2WX_WECHAT_APPID", "wx_env_value")
	os.Setenv("MD2WX_API_KEY", "env_key")
	defer os.Unsetenv("MD2WX_WECHAT_APPID")
	defer os.Unsetenv("MD2WX_API_KEY")

	// 加载配置
	loadedCfg, err := Load()
	if err != nil {
		t.Fatalf("Load() failed: %v", err)
	}

	// 环境变量应该覆盖文件配置
	if loadedCfg.WechatAppID != "wx_env_value" {
		t.Errorf("WechatAppID = %s, want wx_env_value", loadedCfg.WechatAppID)
	}
	if loadedCfg.APIKey != "env_key" {
		t.Errorf("APIKey = %s, want env_key", loadedCfg.APIKey)
	}
	if loadedCfg.WechatAppSecret != "file_secret" {
		t.Errorf("WechatAppSecret = %s, want file_secret", loadedCfg.WechatAppSecret)
	}
}

func TestList(t *testing.T) {
	// 创建临时目录
	tmpDir := t.TempDir()

	// 设置临时主目录
	oldHome := os.Getenv("HOME")
	oldConfigDir := configDir
	oldConfigPath := configPath
	defer func() {
		os.Setenv("HOME", oldHome)
		configDir = oldConfigDir
		configPath = oldConfigPath
	}()
	os.Setenv("HOME", tmpDir)
	configDir = filepath.Join(tmpDir, ConfigDir)
	configPath = filepath.Join(configDir, ConfigFile)

	// 设置一些配置
	if err := Set("api-key", "test_key_12345678"); err != nil {
		t.Fatalf("Set() failed: %v", err)
	}

	// 列出配置
	configs, err := List()
	if err != nil {
		t.Fatalf("List() failed: %v", err)
	}

	if configs["api_key"] == "" {
		t.Error("api_key not found in config list")
	}
	// 检查掩码
	if configs["api_key"] == "test_key_12345678" {
		t.Error("api_key should be masked")
	}
}
