# md2wechat-lite - Agent 工作规范

本文件面向在本仓库内工作的代码智能体（Agent），用于统一开发、测试、文档与发布流程。

## 1. 项目概述

- 项目名：`md2wechat-lite`
- CLI 名：`md2wx`
- 仓库：`https://github.com/geekjourneyx/md2wechat-lite`
- 语言：Go `1.24+`
- CLI 框架：`cobra`

## 2. 项目结构

```text
md2wechat-lite/
├── cli/
│   ├── main.go
│   ├── article-draft.go
│   ├── newspic-draft.go
│   ├── batch-upload.go
│   ├── config.go
│   ├── themes.go
│   ├── pkg/
│   │   ├── api/
│   │   ├── config/
│   │   ├── themes/
│   │   └── output/
│   └── scripts/
├── skills/
│   └── md2wechat-lite/
│       └── SKILL.md
├── AGENTS.md
├── CLAUDE.md
├── CHANGELOG.md
└── README.md
```

## 3. 开发前置认知

1. 本项目是“薄命令层 + API 客户端”架构，核心业务在服务端，CLI 负责参数校验、请求组装、结果输出。
2. 配置文件路径为 `~/.md2wx/config.yaml`，但文件内容是 `key=value` 形式（不是 YAML 结构化对象）。
3. 输出格式以 JSON 为主（`cli/pkg/output`）。

## 4. 图片与封面约束（必须遵守）

1. `article-draft --cover-image`、`newspic-draft --images`、`batch-upload --images` 仅支持公网 URL。
2. 不支持本地图片路径（如 `./a.jpg`）与通配符（如 `*.jpg`）。
3. `article-draft` 不支持通过 stdin 管道直接读取 Markdown（如 `cat a.md | md2wx article-draft ...`）。
4. 为避免服务端 `invalid media_id` 等错误，`article-draft` 建议始终传 `--cover-image`。

## 5. 提交前检查清单

按顺序执行：

1. 编译检查
```bash
go build -o md2wx ./cli
```

2. 单元测试与静态检查
```bash
go test ./cli/pkg/...
go vet ./cli/...
```

3. 文档一致性检查
- 代码变更与 `skills/md2wechat-lite/SKILL.md` 一致。
- `README.md` 示例命令与当前实现一致。
- 涉及图片参数时仅使用 URL 示例，不出现本地路径/通配符误导。

4. SKILL 规范检查（`skills/md2wechat-lite/SKILL.md`）
- 包含 YAML frontmatter（`name` + `description`）。
- `name` 为小写连字符格式，长度小于 64。
- `description` 使用第三人称并覆盖功能与场景。
- 主体内容控制在 500 行内，使用渐进式披露。

## 6. 安全与示例数据规范

1. 文档、帮助文案、安装脚本、测试样例中不得使用疑似真实凭据。
2. 示例统一使用占位值，例如：
- `wx_appid_example`
- `api_key_example`
- `wx_test_appid_example`
3. 不提交真实 `AppID/AppSecret/API Key`。

## 7. Commit 规范

提交格式：

```text
<type>(<scope>): <subject>
```

常用 `type`：
- `feat`
- `fix`
- `docs`
- `style`
- `refactor`
- `test`
- `chore`

## 8. 版本与发版

1. 修改功能或文档后，按需更新 `CHANGELOG.md` 的 `[Unreleased]`。
2. 发版前确认构建、测试、文档一致性全部通过。
3. Tag 与 release 版本号保持一致（示例：`v1.0.x`）。

## 9. 参考链接

- API 文档：`https://md2wechat.app/api-docs`
- 主题预览：`https://md2wechat.app/theme-gallery`
- 项目仓库：`https://github.com/geekjourneyx/md2wechat-lite`
- Skill 仓库：`https://github.com/geekjourneyx/md2wechat-skill`
