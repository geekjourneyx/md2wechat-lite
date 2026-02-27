package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/geekjourneyx/md2wechat-lite/cli/pkg/api"
	"github.com/geekjourneyx/md2wechat-lite/cli/pkg/output"
	"github.com/geekjourneyx/md2wechat-lite/cli/pkg/themes"
	"github.com/spf13/cobra"
)

// ArticleDraftCmd 图文草稿命令
var ArticleDraftCmd = &cobra.Command{
	Use:   "article-draft",
	Short: "创建图文消息草稿",
	Long:  `将 Markdown 内容转换为微信公众号格式并创建图文草稿`,
	Args:  cobra.NoArgs,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return validateArticleDraftFlags()
	},
	Run: runArticleDraft,
}

var (
	flagMarkdown       string
	flagMarkdownFile   string
	flagTheme          string
	flagFontSize       string
	flagBackgroundType string
	flagConvertVersion string
	flagCoverImage     string
)

func init() {
	ArticleDraftCmd.Flags().StringVar(&flagMarkdown, "markdown", "", "Markdown 内容")
	ArticleDraftCmd.Flags().StringVar(&flagMarkdownFile, "file", "", "Markdown 文件路径")
	ArticleDraftCmd.Flags().StringVar(&flagTheme, "theme", "", "主题名称（默认从配置读取）")
	ArticleDraftCmd.Flags().StringVar(&flagFontSize, "font-size", "", "字体大小 (small/medium/large)")
	ArticleDraftCmd.Flags().StringVar(&flagBackgroundType, "background-type", "", "背景类型 (default/grid/none)")
	ArticleDraftCmd.Flags().StringVar(&flagConvertVersion, "convert-version", "v2", "转换版本")
	ArticleDraftCmd.Flags().StringVar(&flagCoverImage, "cover-image", "", "封面图片 URL")
}

func validateArticleDraftFlags() error {
	// 检查 Markdown 来源
	if flagMarkdown == "" && flagMarkdownFile == "" {
		return fmt.Errorf("必须提供 --markdown 或 --file 参数")
	}
	if flagMarkdown != "" && flagMarkdownFile != "" {
		return fmt.Errorf("--markdown 和 --file 不能同时使用")
	}

	// 检查主题是否有效（如果指定了的话）
	if flagTheme != "" && !themes.IsValidTheme(flagTheme) {
		return fmt.Errorf("无效的主题: %s，使用 'themes list' 查看可用主题", flagTheme)
	}

	// 使用配置中的默认主题（如果未指定）
	if flagTheme == "" {
		if cfg.DefaultTheme != "" {
			flagTheme = cfg.DefaultTheme
		} else {
			flagTheme = "default"
		}
	}

	// 使用配置中的默认背景类型（如果未指定）
	if flagBackgroundType == "" {
		if cfg.DefaultBackgroundType != "" {
			flagBackgroundType = cfg.DefaultBackgroundType
		} else {
			flagBackgroundType = "none"
		}
	}

	// 使用配置中的默认字体大小（如果未指定）
	if flagFontSize == "" {
		if cfg.DefaultFontSize != "" {
			flagFontSize = cfg.DefaultFontSize
		} else {
			flagFontSize = "medium"
		}
	}

	// 检查配置
	if cfg.WechatAppID == "" {
		return fmt.Errorf("wechat_appid 未配置，请使用 'config set wechat-appid' 设置")
	}
	if cfg.WechatAppSecret == "" {
		return fmt.Errorf("wechat_appsecret 未配置，请使用 'config set wechat-appsecret' 设置")
	}
	if cfg.APIKey == "" {
		return fmt.Errorf("api_key 未配置，请使用 'config set api-key' 设置")
	}

	return nil
}

func runArticleDraft(cmd *cobra.Command, args []string) {
	// 获取 Markdown 内容
	markdown := flagMarkdown
	if flagMarkdownFile != "" {
		content, err := readFileContent(flagMarkdownFile)
		if err != nil {
			output.Error(err)
		}
		markdown = content
	}

	// 获取 API Base URL（命令行参数优先）
	apiBase := cfg.APIBaseURL
	if apiBaseFlag, _ := cmd.Parent().PersistentFlags().GetString("api-base"); apiBaseFlag != "" {
		apiBase = apiBaseFlag
	}

	// 获取 API Key（命令行参数优先）
	apiKey := cfg.APIKey
	if apiKeyFlag, _ := cmd.Parent().PersistentFlags().GetString("api-key"); apiKeyFlag != "" {
		apiKey = apiKeyFlag
	}

	// 创建 API 客户端
	client := api.NewClient(apiBase, cfg.WechatAppID, cfg.WechatAppSecret, apiKey)

	// 构建请求
	req := &api.ArticleDraftRequest{
		Markdown:       markdown,
		Theme:          flagTheme,
		FontSize:       flagFontSize,
		BackgroundType: flagBackgroundType,
		ConvertVersion: flagConvertVersion,
		CoverImageUrl:  flagCoverImage,
	}

	// 调用 API
	resp, err := client.ArticleDraft(req)
	if err != nil {
		output.Error(err)
	}

	// 输出结果
	result := map[string]interface{}{
		"success": resp.Code == 0,
	}
	if resp.Code == 0 {
		result["draft_id"] = resp.Data.DraftID
		result["media_id"] = resp.Data.MediaID
		result["published"] = resp.Data.Published
		if resp.Data.HTML != "" {
			result["html_preview"] = resp.Data.HTML[:min(200, len(resp.Data.HTML))] + "..."
		}
		output.Success(result)
	} else {
		output.ErrorWithCode(fmt.Sprintf("API_ERROR_%d", resp.Code), resp.Msg)
	}
}

// readFileContent 读取文件内容
func readFileContent(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("读取文件失败: %w", err)
	}
	return string(data), nil
}

// min 返回两个整数中的较小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// parseImageList 解析图片列表
func parseImageList(images string) []string {
	if images == "" {
		return nil
	}
	parts := strings.Split(images, ",")
	result := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			result = append(result, p)
		}
	}
	return result
}
