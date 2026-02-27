package main

import (
	"fmt"

	"github.com/geekjourneyx/md2wechat-lite/cli/pkg/config"
	"github.com/geekjourneyx/md2wechat-lite/cli/pkg/output"
	"github.com/spf13/cobra"
)

// ConfigCmd 配置命令
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "管理配置文件",
	Long:  `管理 md2wechat-lite 的配置文件，支持设置、获取和列出配置项。`,
}

var (
	setKey, setValue string
)

// configSetCmd 设置配置命令
var configSetCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "设置配置项",
	Long:  `设置指定配置项的值。支持: wechat-appid, wechat-appsecret, api-key, api-base`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value := args[1]

		if err := config.Set(key, value); err != nil {
			output.Error(err)
		}

		output.PrintSuccess("✓ 配置已保存: %s = %s", key, maskIfSensitive(key, value))
	},
}

// configGetCmd 获取配置命令
var configGetCmd = &cobra.Command{
	Use:   "get <key>",
	Short: "获取配置项",
	Long:  `获取指定配置项的值`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]

		value, err := config.Get(key)
		if err != nil {
			output.Error(err)
		}

		fmt.Println(value)
	},
}

// configListCmd 列出配置命令
var configListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出所有配置",
	Long:  `列出所有已配置的项`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.List()
		if err != nil {
			output.Error(err)
		}

		if len(cfg) == 0 {
			output.PrintSuccess("暂无配置，请使用 'config set' 命令设置配置")
			return
		}

		output.PrintSuccess("当前配置:")
		for key, value := range cfg {
			fmt.Printf("  %s: %s\n", key, value)
		}
		fmt.Printf("\n配置文件: %s\n", config.GetConfigPath())
	},
}

// configPathCmd 显示配置路径命令
var configPathCmd = &cobra.Command{
	Use:   "path",
	Short: "显示配置文件路径",
	Long:  `显示配置文件的完整路径`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(config.GetConfigPath())
	},
}

func init() {
	ConfigCmd.AddCommand(configSetCmd)
	ConfigCmd.AddCommand(configGetCmd)
	ConfigCmd.AddCommand(configListCmd)
	ConfigCmd.AddCommand(configPathCmd)
}

// maskIfSensitive 如果是敏感信息则掩码
func maskIfSensitive(key, value string) string {
	switch key {
	case "wechat-appid", "wechat_appid",
		"wechat-appsecret", "wechat_appsecret",
		"api-key", "api_key":
		if len(value) <= 8 {
			return "***"
		}
		return value[:4] + "***" + value[len(value)-4:]
	default:
		return value
	}
}
