// Package themes 提供主题列表管理和查询功能。
//
// 支持的主题:
//   - 6 种内置主题: default, bytedance, chinese, apple, sports, cyber
//   - 32 种模板主题: {模板}-{色调} 组合
//     模板: minimal, focus, elegant, bold
//     色调: gold, green, blue, orange, red, navy, gray, sky
package themes

// 6 种内置主题
var BuiltInThemes = []string{
	"default",   // 默认主题 - 微信经典风格
	"bytedance", // 字节范 - 科技现代风格
	"chinese",   // 中国风 - 古典雅致风格
	"apple",     // 苹果范 - 视觉渐变风格
	"sports",    // 运动风 - 活力动感风格
	"cyber",     // 赛博朋克 - 未来科技风格
}

// 32 种模板主题（模板-色调组合）
var TemplateThemes = []string{
	// minimal 模板（纯色文字，无装饰）
	"minimal-gold", "minimal-green", "minimal-blue", "minimal-orange",
	"minimal-red", "minimal-navy", "minimal-gray", "minimal-sky",
	// focus 模板（居中对称，双横线）
	"focus-gold", "focus-green", "focus-blue", "focus-orange",
	"focus-red", "focus-navy", "focus-gray", "focus-sky",
	// elegant 模板（左边框递减，渐变）
	"elegant-gold", "elegant-green", "elegant-blue", "elegant-orange",
	"elegant-red", "elegant-navy", "elegant-gray", "elegant-sky",
	// bold 模板（满底色标题，投影）
	"bold-gold", "bold-green", "bold-blue", "bold-orange",
	"bold-red", "bold-navy", "bold-gray", "bold-sky",
}

// AllThemes 所有主题（内置 + 模板）
var AllThemes = append(append([]string{}, BuiltInThemes...), TemplateThemes...)

// ThemeDescriptions 主题描述映射
var ThemeDescriptions = map[string]string{
	"default":   "微信经典风格",
	"bytedance": "科技现代风格",
	"chinese":   "古典雅致风格",
	"apple":     "视觉渐变风格",
	"sports":    "活力动感风格",
	"cyber":     "未来科技风格",
}

// TemplateColors 模板色调描述
var TemplateColors = map[string]string{
	"gold":   "古铜金 #C8A062",
	"green":  "翡翠绿 #2BAE85",
	"blue":   "宝石蓝 #4B6EF5",
	"orange": "暖阳橙 #F89A3A",
	"red":    "中国红 #F25C54",
	"navy":   "深海蓝 #1F4F8A",
	"gray":   "石墨灰 #4E5969",
	"sky":    "天空蓝 #3A7FD5",
}

// TemplateStyles 模板风格描述
var TemplateStyles = map[string]string{
	"minimal": "简约 - 纯色文字，无装饰",
	"focus":   "聚焦 - 居中对称，双横线",
	"elegant": "精致 - 左边框递减，渐变",
	"bold":    "醒目 - 满底色标题，投影",
}

// IsValidTheme 检查主题是否有效
func IsValidTheme(theme string) bool {
	for _, t := range AllThemes {
		if t == theme {
			return true
		}
	}
	return false
}

// GetThemeDescription 获取主题描述
func GetThemeDescription(theme string) string {
	if desc, ok := ThemeDescriptions[theme]; ok {
		return desc
	}
	// 解析模板主题（如 "minimal-gold"）
	for style, styleDesc := range TemplateStyles {
		for color, colorDesc := range TemplateColors {
			if theme == style+"-"+color {
				return styleDesc + " - " + colorDesc
			}
		}
	}
	return "自定义主题"
}
