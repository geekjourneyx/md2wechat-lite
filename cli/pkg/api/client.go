// Package api 提供 md2wechat API 服务的 HTTP 客户端。
//
// 客户端支持三种主要操作：
//   - 创建图文草稿 (ArticleDraft)
//   - 创建小绿书草稿 (NewspicDraft)
//   - 批量上传素材 (BatchUpload)
//
// 使用方法：
//
//	client := api.NewClient(baseURL, appID, appSecret, apiKey)
//	resp, err := client.ArticleDraft(req)
package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client API 客户端
type Client struct {
	baseURL         string
	wechatAppID     string
	wechatAppSecret string
	apiKey          string
	httpClient      *http.Client
	timeout         time.Duration
}

// NewClient 创建 API 客户端
func NewClient(baseURL, wechatAppID, wechatAppSecret, apiKey string) *Client {
	return &Client{
		baseURL:         baseURL,
		wechatAppID:     wechatAppID,
		wechatAppSecret: wechatAppSecret,
		apiKey:          apiKey,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		timeout: 30 * time.Second,
	}
}

// ArticleDraftRequest 图文草稿请求
type ArticleDraftRequest struct {
	Markdown       string `json:"markdown"`
	Theme          string `json:"theme,omitempty"`
	FontSize       string `json:"fontSize,omitempty"`
	BackgroundType string `json:"backgroundType,omitempty"`
	ConvertVersion string `json:"convertVersion,omitempty"`
	CoverImageUrl  string `json:"coverImageUrl,omitempty"`
}

// NewspicDraftRequest 小绿书草稿请求
type NewspicDraftRequest struct {
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	ImageUrls []string `json:"imageUrls"`
}

// BatchUploadRequest 批量上传请求
type BatchUploadRequest struct {
	ImageUrls []string `json:"imageUrls"`
}

// APIResponse 通用 API 响应
type APIResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

// ArticleDraftResponse 图文草稿响应
type ArticleDraftResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		DraftID   string `json:"draft_id,omitempty"`
		MediaID   string `json:"media_id,omitempty"`
		HTML      string `json:"html,omitempty"`
		Published bool   `json:"published,omitempty"`
	} `json:"data,omitempty"`
}

// NewspicDraftResponse 小绿书草稿响应
type NewspicDraftResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		DraftID   string `json:"draft_id,omitempty"`
		Published bool   `json:"published,omitempty"`
	} `json:"data,omitempty"`
}

// BatchUploadResponse 批量上传响应
type BatchUploadResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Results []UploadResult `json:"results,omitempty"`
	} `json:"data,omitempty"`
}

// UploadResult 上传结果
type UploadResult struct {
	URL     string `json:"url,omitempty"`
	MediaID string `json:"media_id,omitempty"`
	Success bool   `json:"success,omitempty"`
	Error   string `json:"error,omitempty"`
}

// ArticleDraft 创建图文草稿
func (c *Client) ArticleDraft(req *ArticleDraftRequest) (*ArticleDraftResponse, error) {
	endpoint := "/api/v1/article-draft"
	var resp ArticleDraftResponse
	if err := c.doRequest("POST", endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// NewspicDraft 创建小绿书草稿
func (c *Client) NewspicDraft(req *NewspicDraftRequest) (*NewspicDraftResponse, error) {
	endpoint := "/api/v1/newspic-draft"
	var resp NewspicDraftResponse
	if err := c.doRequest("POST", endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// BatchUpload 批量上传素材
func (c *Client) BatchUpload(req *BatchUploadRequest) (*BatchUploadResponse, error) {
	endpoint := "/api/v1/batch-upload"
	var resp BatchUploadResponse
	if err := c.doRequest("POST", endpoint, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// doRequest 执行 HTTP 请求
func (c *Client) doRequest(method, endpoint string, body any, resp any) error {
	// 构建完整 URL
	url := c.baseURL + endpoint

	// 序列化请求体
	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("序列化请求失败: %w", err)
		}
		reqBody = bytes.NewReader(jsonData)
	}

	// 创建请求
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return fmt.Errorf("创建请求失败: %w", err)
	}

	// 设置认证 Headers
	c.setAuthHeaders(req)

	// 发送请求
	httpResp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("请求失败: %w", err)
	}
	defer httpResp.Body.Close()

	// 读取响应
	respData, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return fmt.Errorf("读取响应失败: %w", err)
	}

	// 检查 HTTP 状态码
	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP 错误: %d, 响应: %s", httpResp.StatusCode, string(respData))
	}

	// 解析响应
	if err := json.Unmarshal(respData, resp); err != nil {
		return fmt.Errorf("解析响应失败: %w (响应: %s)", err, string(respData))
	}

	// 检查业务状态码
	if apiResp, ok := resp.(*APIResponse); ok && apiResp.Code != 0 {
		return fmt.Errorf("API 错误 (code %d): %s", apiResp.Code, apiResp.Msg)
	}

	// 检查响应状态码（针对有 Code 字段的响应）
	checkCode := func(code int, msg string) {
		if code != 0 {
			// 非零 code 表示业务错误，但不返回 error，让上层处理
		}
	}

	switch r := resp.(type) {
	case *ArticleDraftResponse:
		checkCode(r.Code, r.Msg)
	case *NewspicDraftResponse:
		checkCode(r.Code, r.Msg)
	case *BatchUploadResponse:
		checkCode(r.Code, r.Msg)
	}

	return nil
}

// setAuthHeaders 设置认证 Headers
func (c *Client) setAuthHeaders(req *http.Request) {
	req.Header.Set("Wechat-Appid", c.wechatAppID)
	req.Header.Set("Wechat-App-Secret", c.wechatAppSecret)
	req.Header.Set("Md2wechat-API-Key", c.apiKey)
	req.Header.Set("Content-Type", "application/json")
}

// SetTimeout 设置请求超时
func (c *Client) SetTimeout(timeout time.Duration) {
	c.timeout = timeout
	c.httpClient.Timeout = timeout
}
