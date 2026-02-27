package main

import (
	"fmt"

	"github.com/geekjourneyx/md2wechat-lite/cli/pkg/api"
	"github.com/geekjourneyx/md2wechat-lite/cli/pkg/output"
	"github.com/spf13/cobra"
)

// NewspicDraftCmd 小绿书草稿命令
var NewspicDraftCmd = &cobra.Command{
	Use:   "newspic-draft",
	Short: "创建小绿书草稿",
	Long:  `创建微信公众号小绿书（图片文章）草稿`,
	Args:  cobra.NoArgs,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return validateNewspicDraftFlags()
	},
	Run: runNewspicDraft,
}

var (
	flagTitle       string
	flagContent     string
	flagImages      string
	flagContentFile string
)

func init() {
	NewspicDraftCmd.Flags().StringVar(&flagTitle, "title", "", "文章标题")
	NewspicDraftCmd.Flags().StringVar(&flagContent, "content", "", "正文内容")
	NewspicDraftCmd.Flags().StringVar(&flagImages, "images", "", "图片 URL，多个用逗号分隔")
	NewspicDraftCmd.Flags().StringVar(&flagContentFile, "content-file", "", "正文内容文件路径")
}

func validateNewspicDraftFlags() error {
	if flagTitle == "" {
		return fmt.Errorf("必须提供 --title 参数")
	}

	if flagContent == "" && flagContentFile == "" {
		return fmt.Errorf("必须提供 --content 或 --content-file 参数")
	}

	if flagContent != "" && flagContentFile != "" {
		return fmt.Errorf("--content 和 --content-file 不能同时使用")
	}

	if flagImages == "" {
		return fmt.Errorf("必须提供 --images 参数（至少一张图片）")
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

func runNewspicDraft(cmd *cobra.Command, args []string) {
	// 获取内容
	content := flagContent
	if flagContentFile != "" {
		c, err := readFileContent(flagContentFile)
		if err != nil {
			output.Error(err)
		}
		content = c
	}

	// 解析图片列表
	imageUrls := parseImageList(flagImages)

	if len(imageUrls) == 0 {
		output.Error(fmt.Errorf("至少需要一张图片"))
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
	req := &api.NewspicDraftRequest{
		Title:     flagTitle,
		Content:   content,
		ImageUrls: imageUrls,
	}

	// 调用 API
	resp, err := client.NewspicDraft(req)
	if err != nil {
		output.Error(err)
	}

	// 输出结果
	if resp.Code == 0 {
		output.Success(map[string]interface{}{
			"draft_id":  resp.Data.DraftID,
			"published": resp.Data.Published,
		})
	} else {
		output.ErrorWithCode(fmt.Sprintf("API_ERROR_%d", resp.Code), resp.Msg)
	}
}
