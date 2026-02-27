// Package output 提供统一的输出格式化功能。
//
// 所有输出为 JSON 格式，包含 success 字段表示操作是否成功。
// 成功时包含 data 字段，失败时包含 error 字段。
package output

import (
	"encoding/json"
	"fmt"
	"os"
)

// SuccessResponse 成功响应
type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorResponse 错误响应
type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	Code    string `json:"code,omitempty"`
}

// Success 输出成功响应
func Success(data interface{}) {
	resp := SuccessResponse{
		Success: true,
		Data:    data,
	}
	printJSON(resp)
}

// Error 输出错误响应
func Error(err error) {
	resp := ErrorResponse{
		Success: false,
		Error:   err.Error(),
	}
	printJSON(resp)
	os.Exit(1)
}

// ErrorWithCode 输出带错误码的错误响应
func ErrorWithCode(code, message string) {
	resp := ErrorResponse{
		Success: false,
		Error:   message,
		Code:    code,
	}
	printJSON(resp)
	os.Exit(1)
}

// printJSON 打印 JSON
func printJSON(v interface{}) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(v); err != nil {
		fmt.Fprintf(os.Stderr, "JSON 编码错误: %v\n", err)
		os.Exit(1)
	}
}

// PrintSuccess 打印成功消息（兼容函数）
func PrintSuccess(format string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, format+"\n", args...)
}

// PrintError 打印错误消息到 stderr（兼容函数）
func PrintError(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
}
