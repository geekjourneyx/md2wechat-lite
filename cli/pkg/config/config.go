// Package config 提供配置文件管理功能。
//
// 配置文件位置: ~/.md2wx/config.yaml
//
// 支持的配置项:
//   - wechat_appid: 微信 AppID
//   - wechat_appsecret: 微信 AppSecret
//   - api_key: Md2wechat API Key
//   - api_base_url: API 基础 URL
//   - default_theme: 默认主题名称
//   - background_type: 默认背景类型
//   - font_size: 默认字体大小
//
// 配置优先级: 环境变量 > 配置文件 > 默认值
package config

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
)

// Config 应用配置
type Config struct {
	WechatAppID         string `yaml:"wechat_appid" json:"wechat_appid"`
	WechatAppSecret   string `yaml:"wechat_appsecret" json:"wechat_appsecret"`
	APIKey              string `yaml:"api_key" json:"api_key"`
	APIBaseURL          string `yaml:"api_base_url" json:"api_base_url"`
	DefaultTheme       string `yaml:"default_theme" json:"default_theme"`
	DefaultBackgroundType string `yaml:"background_type" json:"background_type"`
	DefaultFontSize     string `yaml:"font_size" json:"font_size"`
}

const (
	// DefaultAPIBaseURL 默认 API 基础 URL
	DefaultAPIBaseURL = "http://111.231.20.31:8080"
	// ConfigDir 配置目录名
	ConfigDir = ".md2wx"
	// ConfigFile 配置文件名
	ConfigFile = "config.yaml"
)

var (
	// configPath 配置文件完整路径
	configPath string
	// configDir 配置目录路径
	configDir string
)

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return
	}
	configDir = filepath.Join(homeDir, ConfigDir)
	configPath = filepath.Join(configDir, ConfigFile)
}

// GetConfigPath 获取配置文件路径
func GetConfigPath() string {
	return configPath
}

// GetConfigDir 获取配置目录路径
func GetConfigDir() string {
	return configDir
}

// Load 从配置文件加载配置
func Load() (*Config, error) {
	cfg := &Config{
		APIBaseURL: DefaultAPIBaseURL,
	}

	// 如果配置文件不存在，返回默认配置
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return cfg, nil
	}

	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 简单的 key=value 解析（为了保持轻量，不依赖 yaml 库）
	// 格式: key=value
	lines := splitLines(data)
	for _, line := range lines {
		line = trimSpace(line)
		if line == "" || line[0] == '#' {
			continue
		}
		key, value, ok := parseKeyValue(line)
		if !ok {
			continue
		}
		switch key {
		case "wechat_appid":
			cfg.WechatAppID = value
		case "wechat_appsecret":
			cfg.WechatAppSecret = value
		case "api_key":
			cfg.APIKey = value
		case "api_base_url":
			cfg.APIBaseURL = value
		case "default_theme":
			cfg.DefaultTheme = value
		case "background_type":
			cfg.DefaultBackgroundType = value
		case "font_size":
			cfg.DefaultFontSize = value
		}
	}

	// 环境变量覆盖（优先级更高）
	if v := os.Getenv("MD2WX_WECHAT_APPID"); v != "" {
		cfg.WechatAppID = v
	}
	if v := os.Getenv("MD2WX_WECHAT_APPSECRET"); v != "" {
		cfg.WechatAppSecret = v
	}
	if v := os.Getenv("MD2WX_API_KEY"); v != "" {
		cfg.APIKey = v
	}
	if v := os.Getenv("MD2WX_API_BASE_URL"); v != "" {
		cfg.APIBaseURL = v
	}
	if v := os.Getenv("MD2WX_DEFAULT_THEME"); v != "" {
		cfg.DefaultTheme = v
	}
	if v := os.Getenv("MD2WX_BACKGROUND_TYPE"); v != "" {
		cfg.DefaultBackgroundType = v
	}
	if v := os.Getenv("MD2WX_FONT_SIZE"); v != "" {
		cfg.DefaultFontSize = v
	}

	return cfg, nil
}

// Save 保存配置到文件
func Save(cfg *Config) error {
	// 确保配置目录存在
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("创建配置目录失败: %w", err)
	}

	// 构建配置内容
	var content string
	content += "# md2wx 配置文件\n"
	content += "# 可通过环境变量覆盖（优先级更高）\n"
	content += "#\n"
	content += "# 配置项说明：\n"
	content += "#   wechat_appid    - 微信公众号 AppID（必填）\n"
	content += "#   wechat_appsecret - 微信公众号 AppSecret（必填）\n"
	content += "#   api_key         - md2wx API Key（必填，获取地址：https://www.md2wechat.cn/api-docs）\n"
	content += "#   api_base_url    - API 基础 URL（可选，默认：http://111.231.20.31:8080）\n"
	content += "#   default_theme   - 默认主题（可选）\n"
	content += "#     内置主题: default, bytedance, chinese, apple, sports, cyber\n"
	content += "#     模板主题: minimal-gold, focus-blue, elegant-red, bold-navy 等 32 种\n"
	content += "#   background_type - 默认背景类型（可选）\n"
	content += "#     可选值: none（无）, default（默认）, grid（网格）\n"
	content += "#   font_size       - 默认字体大小（可选）\n"
	content += "#     可选值: small（小）, medium（中）, large（大）\n"
	content += "#\n\n"

	if cfg.WechatAppID != "" {
		content += fmt.Sprintf("wechat_appid=%s\n", cfg.WechatAppID)
	}
	if cfg.WechatAppSecret != "" {
		content += fmt.Sprintf("wechat_appsecret=%s\n", cfg.WechatAppSecret)
	}
	if cfg.APIKey != "" {
		content += fmt.Sprintf("api_key=%s\n", cfg.APIKey)
	}
	if cfg.APIBaseURL != "" && cfg.APIBaseURL != DefaultAPIBaseURL {
		content += fmt.Sprintf("api_base_url=%s\n", cfg.APIBaseURL)
	}
	if cfg.DefaultTheme != "" && cfg.DefaultTheme != "default" {
		content += fmt.Sprintf("default_theme=%s\n", cfg.DefaultTheme)
	}
	if cfg.DefaultBackgroundType != "" && cfg.DefaultBackgroundType != "none" {
		content += fmt.Sprintf("background_type=%s\n", cfg.DefaultBackgroundType)
	}
	if cfg.DefaultFontSize != "" && cfg.DefaultFontSize != "medium" {
		content += fmt.Sprintf("font_size=%s\n", cfg.DefaultFontSize)
	}

	// 写入文件
	if err := os.WriteFile(configPath, []byte(content), 0600); err != nil {
		return fmt.Errorf("写入配置文件失败: %w", err)
	}

	return nil
}

// Set 设置单个配置项
func Set(key, value string) error {
	cfg, err := Load()
	if err != nil {
		return err
	}

	switch key {
	case "wechat-appid", "wechat_appid":
		cfg.WechatAppID = value
	case "wechat-appsecret", "wechat_appsecret":
		cfg.WechatAppSecret = value
	case "api-key", "api_key":
		cfg.APIKey = value
	case "api-base", "api_base_url":
		cfg.APIBaseURL = value
	case "default-theme", "default_theme":
		cfg.DefaultTheme = value
	case "background-type", "background_type":
		cfg.DefaultBackgroundType = value
	case "font-size", "font_size":
		cfg.DefaultFontSize = value
	default:
		return fmt.Errorf("未知的配置项: %s", key)
	}

	return Save(cfg)
}

// Get 获取单个配置项
func Get(key string) (string, error) {
	cfg, err := Load()
	if err != nil {
		return "", err
	}

	switch key {
	case "wechat-appid", "wechat_appid":
		if cfg.WechatAppID == "" {
			return "", fmt.Errorf("wechat_appid 未配置")
		}
		return cfg.WechatAppID, nil
	case "wechat-appsecret", "wechat_appsecret":
		if cfg.WechatAppSecret == "" {
			return "", fmt.Errorf("wechat_appsecret 未配置")
		}
		return cfg.WechatAppSecret, nil
	case "api-key", "api_key":
		if cfg.APIKey == "" {
			return "", fmt.Errorf("api_key 未配置")
		}
		return cfg.APIKey, nil
	case "api-base", "api_base_url":
		return cfg.APIBaseURL, nil
	case "default-theme", "default_theme":
		if cfg.DefaultTheme == "" {
			return "default", nil
		}
		return cfg.DefaultTheme, nil
	case "background-type", "background_type":
		if cfg.DefaultBackgroundType == "" {
			return "none", nil
		}
		return cfg.DefaultBackgroundType, nil
	case "font-size", "font_size":
		if cfg.DefaultFontSize == "" {
			return "medium", nil
		}
		return cfg.DefaultFontSize, nil
	default:
		return "", fmt.Errorf("未知的配置项: %s", key)
	}
}

// List 列出所有配置项
func List() (map[string]string, error) {
	cfg, err := Load()
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)
	if cfg.WechatAppID != "" {
		result["wechat_appid"] = maskSensitive(cfg.WechatAppID)
	}
	if cfg.WechatAppSecret != "" {
		result["wechat_appsecret"] = maskSensitive(cfg.WechatAppSecret)
	}
	if cfg.APIKey != "" {
		result["api_key"] = maskSensitive(cfg.APIKey)
	}
	result["api_base_url"] = cfg.APIBaseURL
	if cfg.DefaultTheme != "" {
		result["default_theme"] = cfg.DefaultTheme
	} else {
		result["default_theme"] = "default"
	}
	if cfg.DefaultBackgroundType != "" {
		result["background_type"] = cfg.DefaultBackgroundType
	} else {
		result["background_type"] = "none"
	}
	if cfg.DefaultFontSize != "" {
		result["font_size"] = cfg.DefaultFontSize
	} else {
		result["font_size"] = "medium"
	}

	return result, nil
}

// maskSensitive 掩码敏感信息
func maskSensitive(s string) string {
	if len(s) <= 8 {
		return "***"
	}
	return s[:4] + "***" + s[len(s)-4:]
}

// 简单的字符串处理辅助函数
func splitLines(data []byte) []string {
	var lines []string
	start := 0
	for i, b := range data {
		if b == '\n' {
			lines = append(lines, string(data[start:i]))
			start = i + 1
		}
	}
	if start < len(data) {
		lines = append(lines, string(data[start:]))
	}
	return lines
}

func trimSpace(s string) string {
	start := 0
	for start < len(s) && (s[start] == ' ' || s[start] == '\t' || s[start] == '\r') {
		start++
	}
	end := len(s)
	for end > start && (s[end-1] == ' ' || s[end-1] == '\t' || s[end-1] == '\r') {
		end--
	}
	return s[start:end]
}

func parseKeyValue(line string) (string, string, bool) {
	idx := slices.Index([]byte(line), '=')
	if idx == -1 {
		return "", "", false
	}
	key := trimSpace(line[:idx])
	value := trimSpace(line[idx+1:])
	return key, value, true
}
