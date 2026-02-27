package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	client := NewClient("http://example.com", "appid", "secret", "key")

	if client.baseURL != "http://example.com" {
		t.Errorf("baseURL = %s, want http://example.com", client.baseURL)
	}
	if client.wechatAppID != "appid" {
		t.Errorf("wechatAppID = %s, want appid", client.wechatAppID)
	}
	if client.wechatAppSecret != "secret" {
		t.Errorf("wechatAppSecret = %s, want secret", client.wechatAppSecret)
	}
	if client.apiKey != "key" {
		t.Errorf("apiKey = %s, want key", client.apiKey)
	}
	if client.httpClient == nil {
		t.Error("httpClient is nil")
	}
}

func TestSetAuthHeaders(t *testing.T) {
	client := NewClient("http://example.com", "test-appid", "test-secret", "test-key")

	req, _ := http.NewRequest("POST", "http://example.com", nil)
	client.setAuthHeaders(req)

	if req.Header.Get("Wechat-Appid") != "test-appid" {
		t.Errorf("Wechat-Appid = %s, want test-appid", req.Header.Get("Wechat-Appid"))
	}
	if req.Header.Get("Wechat-App-Secret") != "test-secret" {
		t.Errorf("Wechat-App-Secret = %s, want test-secret", req.Header.Get("Wechat-App-Secret"))
	}
	if req.Header.Get("Md2wechat-API-Key") != "test-key" {
		t.Errorf("Md2wechat-API-Key = %s, want test-key", req.Header.Get("Md2wechat-API-Key"))
	}
	if req.Header.Get("Content-Type") != "application/json" {
		t.Errorf("Content-Type = %s, want application/json", req.Header.Get("Content-Type"))
	}
}

func TestArticleDraft_Success(t *testing.T) {
	// 创建 mock 服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 验证请求
		if r.Method != "POST" {
			t.Errorf("Method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/api/v1/article-draft" {
			t.Errorf("Path = %s, want /api/v1/article-draft", r.URL.Path)
		}
		if r.Header.Get("Wechat-Appid") != "test-appid" {
			t.Errorf("Wechat-Appid header missing")
		}

		// 返回成功响应
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code": 0,
			"msg":  "success",
			"data": map[string]interface{}{
				"draft_id":  "draft_123",
				"media_id":  "media_456",
				"published": false,
			},
		})
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-appid", "test-secret", "test-key")

	req := &ArticleDraftRequest{
		Markdown: "# Test\n\nContent",
		Theme:    "default",
	}

	resp, err := client.ArticleDraft(req)
	if err != nil {
		t.Fatalf("ArticleDraft() failed: %v", err)
	}

	if resp.Code != 0 {
		t.Errorf("Code = %d, want 0", resp.Code)
	}
	if resp.Data.DraftID != "draft_123" {
		t.Errorf("DraftID = %s, want draft_123", resp.Data.DraftID)
	}
	if resp.Data.Published != false {
		t.Errorf("Published = %v, want false", resp.Data.Published)
	}
}

func TestArticleDraft_APIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code": 40001,
			"msg":  "invalid api key",
		})
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-appid", "test-secret", "test-key")

	req := &ArticleDraftRequest{
		Markdown: "# Test",
	}

	resp, err := client.ArticleDraft(req)
	// API 错误不会返回 error，而是设置 Code
	if err != nil {
		t.Fatalf("ArticleDraft() failed: %v", err)
	}

	if resp.Code != 40001 {
		t.Errorf("Code = %d, want 40001", resp.Code)
	}
	if resp.Msg != "invalid api key" {
		t.Errorf("Msg = %s, want 'invalid api key'", resp.Msg)
	}
}

func TestNewspicDraft_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/newspic-draft" {
			t.Errorf("Path = %s, want /api/v1/newspic-draft", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code": 0,
			"msg":  "success",
			"data": map[string]interface{}{
				"draft_id":  "newspic_123",
				"published": true,
			},
		})
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-appid", "test-secret", "test-key")

	req := &NewspicDraftRequest{
		Title:     "Test Title",
		Content:   "Test Content",
		ImageUrls: []string{"http://example.com/1.jpg"},
	}

	resp, err := client.NewspicDraft(req)
	if err != nil {
		t.Fatalf("NewspicDraft() failed: %v", err)
	}

	if resp.Code != 0 {
		t.Errorf("Code = %d, want 0", resp.Code)
	}
	if resp.Data.DraftID != "newspic_123" {
		t.Errorf("DraftID = %s, want newspic_123", resp.Data.DraftID)
	}
}

func TestBatchUpload_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/batch-upload" {
			t.Errorf("Path = %s, want /api/v1/batch-upload", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code": 0,
			"msg":  "success",
			"data": map[string]interface{}{
				"results": []map[string]interface{}{
					{"url": "http://example.com/1.jpg", "media_id": "media_1", "success": true},
				},
			},
		})
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-appid", "test-secret", "test-key")

	req := &BatchUploadRequest{
		ImageUrls: []string{"http://example.com/1.jpg"},
	}

	resp, err := client.BatchUpload(req)
	if err != nil {
		t.Fatalf("BatchUpload() failed: %v", err)
	}

	if resp.Code != 0 {
		t.Errorf("Code = %d, want 0", resp.Code)
	}
	if len(resp.Data.Results) != 1 {
		t.Errorf("len(Results) = %d, want 1", len(resp.Data.Results))
	}
}

func TestSetTimeout(t *testing.T) {
	client := NewClient("http://example.com", "appid", "secret", "key")

	client.SetTimeout(60 * time.Second)

	if client.timeout != 60*time.Second {
		t.Errorf("timeout = %v, want 1m0s", client.timeout)
	}
	if client.httpClient.Timeout != 60*time.Second {
		t.Errorf("httpClient.Timeout = %v, want 1m0s", client.httpClient.Timeout)
	}
}
