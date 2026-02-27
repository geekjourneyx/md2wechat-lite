package themes

import (
	"testing"
)

func TestIsValidTheme(t *testing.T) {
	tests := []struct {
		name  string
		theme string
		want  bool
	}{
		{"内置主题 - default", "default", true},
		{"内置主题 - bytedance", "bytedance", true},
		{"内置主题 - cyber", "cyber", true},
		{"模板主题 - minimal-gold", "minimal-gold", true},
		{"模板主题 - bold-navy", "bold-navy", true},
		{"无效主题", "invalid-theme", false},
		{"空字符串", "", false},
		{"部分匹配", "minimal", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidTheme(tt.theme)
			if got != tt.want {
				t.Errorf("IsValidTheme(%q) = %v, want %v", tt.theme, got, tt.want)
			}
		})
	}
}

func TestGetThemeDescription(t *testing.T) {
	tests := []struct {
		name  string
		theme string
		want  string
	}{
		{"内置主题", "default", "微信经典风格"},
		{"内置主题", "bytedance", "科技现代风格"},
		{"内置主题", "chinese", "古典雅致风格"},
		{"模板主题", "minimal-gold", "简约 - 纯色文字，无装饰 - 古铜金 #C8A062"},
		{"无效主题", "invalid", "自定义主题"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetThemeDescription(tt.theme)
			if got != tt.want {
				t.Errorf("GetThemeDescription(%q) = %q, want %q", tt.theme, got, tt.want)
			}
		})
	}
}

func TestAllThemes_Count(t *testing.T) {
	// 6 内置 + 32 模板 = 38
	want := 38
	got := len(AllThemes)

	if got != want {
		t.Errorf("len(AllThemes) = %d, want %d", got, want)
	}
}

func TestAllThemes_NoDuplicates(t *testing.T) {
	seen := make(map[string]bool)
	for _, theme := range AllThemes {
		if seen[theme] {
			t.Errorf("Duplicate theme found: %s", theme)
		}
		seen[theme] = true
	}
}

func TestBuiltInThemes_Count(t *testing.T) {
	want := 6
	got := len(BuiltInThemes)

	if got != want {
		t.Errorf("len(BuiltInThemes) = %d, want %d", got, want)
	}
}

func TestTemplateThemes_Count(t *testing.T) {
	want := 32
	got := len(TemplateThemes)

	if got != want {
		t.Errorf("len(TemplateThemes) = %d, want %d", got, want)
	}
}

func TestThemeDescriptions_Completeness(t *testing.T) {
	// 确保所有内置主题都有描述
	for _, theme := range BuiltInThemes {
		if desc, ok := ThemeDescriptions[theme]; !ok || desc == "" {
			t.Errorf("BuiltIn theme %q missing description", theme)
		}
	}
}

func TestTemplateColors_Completeness(t *testing.T) {
	want := 8
	got := len(TemplateColors)

	if got != want {
		t.Errorf("len(TemplateColors) = %d, want %d", got, want)
	}
}

func TestTemplateStyles_Completeness(t *testing.T) {
	want := 4
	got := len(TemplateStyles)

	if got != want {
		t.Errorf("len(TemplateStyles) = %d, want %d", got, want)
	}
}
