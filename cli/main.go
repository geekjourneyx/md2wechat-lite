package main

import (
	"fmt"
	"os"

	"github.com/geekjourneyx/md2wechat-lite/cli/pkg/config"
	"github.com/geekjourneyx/md2wechat-lite/cli/pkg/output"
	"github.com/spf13/cobra"
)

// 版本信息（构建时注入）
var (
	version   = "1.0.0"
	buildDate = "unknown"
	gitCommit = "unknown"
)

func main() {
	// 执行根命令
	if err := rootCmd.Execute(); err != nil {
		output.Error(err)
		os.Exit(1)
	}
}

// rootCmd 根命令
var rootCmd = &cobra.Command{
	Use:   "md2wx",
	Short: "Markdown 转微信公众号格式 - 轻量级 CLI 工具",
	Long: `md2wechat-lite 是一个轻量级的命令行工具，通过 API 服务将 Markdown
转换为微信公众号格式，支持图文草稿和小绿书创建。

快速开始:
  md2wechat-lite config set wechat-appid "wx_appid_example"
  md2wechat-lite config set api-key "api_key_example"
  md2wechat-lite article-draft --markdown "# 标题\n\n内容" --theme "default"

获取帮助:
  md2wechat-lite --help
  md2wechat-lite [command] --help`,
	SilenceErrors: true,
	SilenceUsage:  true,
}

var cfg *config.Config

// 执行命令前的初始化
func initConfig(cmd *cobra.Command, args []string) error {
	var err error
	cfg, err = config.Load()
	if err != nil {
		return fmt.Errorf("加载配置失败: %w", err)
	}
	return nil
}

func init() {
	// 添加版本标志
	rootCmd.Version = fmt.Sprintf("%s (构建时间: %s, 提交: %s)", version, buildDate, gitCommit)

	// 添加子命令
	rootCmd.AddCommand(ConfigCmd)
	rootCmd.AddCommand(ArticleDraftCmd)
	rootCmd.AddCommand(NewspicDraftCmd)
	rootCmd.AddCommand(BatchUploadCmd)
	rootCmd.AddCommand(ThemesCmd)

	// 持久化标志
	rootCmd.PersistentFlags().StringP("api-base", "a", "", "API 基础 URL (覆盖配置文件)")
	rootCmd.PersistentFlags().StringP("api-key", "k", "", "API Key (覆盖配置文件)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "详细输出")

	// 绑定持久化标志到配置
	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		// 某些命令不需要配置（如 help, version, config set, themes list）
		skipConfig := cmd.Name() == "help" || cmd.Name() == "config" || cmd.Name() == "themes"
		if !skipConfig {
			if err := initConfig(cmd, args); err != nil {
				return err
			}
		}
		return nil
	}
}
