package main

import (
	"fmt"

	"github.com/geekjourneyx/md2wechat-lite/cli/pkg/api"
	"github.com/geekjourneyx/md2wechat-lite/cli/pkg/output"
	"github.com/spf13/cobra"
)

// BatchUploadCmd 批量上传命令
var BatchUploadCmd = &cobra.Command{
	Use:   "batch-upload",
	Short: "批量上传图片素材",
	Long:  `批量上传图片到微信公众号素材库`,
	Args:  cobra.NoArgs,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return validateBatchUploadFlags()
	},
	Run: runBatchUpload,
}

var flagUploadImages string

func init() {
	BatchUploadCmd.Flags().StringVar(&flagUploadImages, "images", "", "图片 URL，多个用逗号分隔")
}

func validateBatchUploadFlags() error {
	if flagUploadImages == "" {
		return fmt.Errorf("必须提供 --images 参数")
	}

	imageUrls := parseImageList(flagUploadImages)
	if len(imageUrls) == 0 {
		return fmt.Errorf("至少需要一张图片")
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

func runBatchUpload(cmd *cobra.Command, args []string) {
	// 解析图片列表
	imageUrls := parseImageList(flagUploadImages)

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
	req := &api.BatchUploadRequest{
		ImageUrls: imageUrls,
	}

	// 调用 API
	resp, err := client.BatchUpload(req)
	if err != nil {
		output.Error(err)
	}

	// 输出结果
	if resp.Code == 0 {
		output.Success(map[string]interface{}{
			"results": resp.Data.Results,
		})
	} else {
		output.ErrorWithCode(fmt.Sprintf("API_ERROR_%d", resp.Code), resp.Msg)
	}
}
