package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/geekjourneyx/md2wechat-lite/cli/pkg/output"
	"github.com/geekjourneyx/md2wechat-lite/cli/pkg/themes"
	"github.com/spf13/cobra"
)

// ThemesCmd 主题命令
var ThemesCmd = &cobra.Command{
	Use:   "themes",
	Short: "管理主题",
	Long:  `查看和管理可用的排版主题`,
}

// themesListCmd 列出主题命令
var themesListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出所有可用主题",
	Long:  `列出所有可用的排版主题（内置主题 + 模板主题）`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		runThemesList()
	},
}

var (
	flagThemesVerbose bool
	flagThemesSearch  string
)

func init() {
	ThemesCmd.AddCommand(themesListCmd)
	themesListCmd.Flags().BoolVarP(&flagThemesVerbose, "verbose", "v", false, "显示详细信息")
	themesListCmd.Flags().StringVarP(&flagThemesSearch, "search", "s", "", "搜索主题")
}

func runThemesList() {
	var themeList []string

	if flagThemesSearch != "" {
		// 搜索主题
		search := strings.ToLower(flagThemesSearch)
		for _, t := range themes.AllThemes {
			if strings.Contains(strings.ToLower(t), search) {
				themeList = append(themeList, t)
			}
		}
	} else {
		// 显示所有主题
		themeList = themes.AllThemes
	}

	if len(themeList) == 0 {
		output.PrintSuccess("未找到匹配的主题")
		return
	}

	sort.Strings(themeList)

	// 输出
	output.PrintSuccess("可用主题 (%d 个):", len(themeList))
	fmt.Println()

	// 内置主题
	fmt.Println("内置主题:")
	for _, t := range themes.BuiltInThemes {
		if contains(themeList, t) {
			if flagThemesVerbose {
				fmt.Printf("  %-20s %s\n", t, themes.GetThemeDescription(t))
			} else {
				fmt.Printf("  %s\n", t)
			}
		}
	}

	// 模板主题
	fmt.Println("\n模板主题 (模板-色调):")
	// 按模板分组
	templates := []string{"minimal", "focus", "elegant", "bold"}
	colors := []string{"gold", "green", "blue", "orange", "red", "navy", "gray", "sky"}

	for _, tmpl := range templates {
		fmt.Printf("  %s:", tmpl)
		if flagThemesVerbose {
			fmt.Printf(" %s\n", themes.TemplateStyles[tmpl])
		} else {
			fmt.Println()
		}
		for _, color := range colors {
			themeName := tmpl + "-" + color
			if contains(themeList, themeName) {
				if flagThemesVerbose {
					fmt.Printf("    %-20s - %s\n", themeName, themes.TemplateColors[color])
				} else {
					fmt.Printf("    %s\n", themeName)
				}
			}
		}
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
