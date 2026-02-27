package output

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"
)

func TestPrintSuccess(t *testing.T) {
	// 捕获 stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Success(map[string]string{"key": "value"})

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)

	// 验证输出
	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	if result["success"] != true {
		t.Errorf("success = %v, want true", result["success"])
	}
	if result["data"] == nil {
		t.Error("data is nil")
	}
}

func TestPrintSuccess_NoData(t *testing.T) {
	// 捕获 stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Success(nil)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)

	// 验证输出
	var result map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	if result["success"] != true {
		t.Errorf("success = %v, want true", result["success"])
	}
	// nil 数据应该被省略
	if _, ok := result["data"]; ok {
		t.Error("data should be omitted when nil")
	}
}

func TestPrintSuccessMessage(t *testing.T) {
	// 捕获 stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintSuccess("test message: %s", "value")

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)

	output := buf.String()
	if output != "test message: value\n" {
		t.Errorf("output = %q, want 'test message: value\\n'", output)
	}
}

func TestPrintErrorMessage(t *testing.T) {
	// 捕获 stderr
	old := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	PrintError("error message: %s", "value")

	w.Close()
	os.Stderr = old

	var buf bytes.Buffer
	buf.ReadFrom(r)

	output := buf.String()
	if output != "error message: value\n" {
		t.Errorf("output = %q, want 'error message: value\\n'", output)
	}
}
