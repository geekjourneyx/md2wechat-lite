# md2wechat-lite - 项目记忆

## 项目概述

**md2wechat-lite** 是一个轻量级的 CLI 工具，通过 API 服务将 Markdown 转换为微信公众号格式。

- **项目名**: md2wechat-lite
- **CLI 名**: md2wx
- **仓库**: https://github.com/geekjourneyx/md2wechat-lite
- **语言**: Go 1.24+
- **框架**: Cobra

## 项目结构

```
md2wechat-lite/
├── cli/                    # CLI 源码
│   ├── main.go             # 程序入口
│   ├── article-draft.go    # 图文草稿命令
│   ├── newspic-draft.go    # 小绿书草稿命令
│   ├── batch-upload.go     # 批量上传命令
│   ├── config.go           # 配置管理命令
│   ├── themes.go           # 主题列表命令
│   ├── pkg/                # 内部包
│   │   ├── api/           # API 客户端
│   │   ├── config/        # 配置文件管理
│   │   ├── themes/        # 主题列表
│   │   └── output/        # 输出格式化
│   └── scripts/           # 安装脚本
├── skills/                 # Claude Code Skill 定义
│   └── md2wechat-lite/
│       └── SKILL.md        # Skill 定义文件
├── CLAUDE.md              # 本文件
├── CHANGELOG.md           # 变更日志
└── README.md              # 项目说明
```

## 提交前工作流程

每次代码提交前，必须按顺序执行以下步骤：

### 1. 编译检查

```bash
cd /root/md2wechat-lite
go build -o md2wx ./cli
```

确保编译无错误。

### 2. 代码检查

```bash
go test ./cli/pkg/...
go vet ./cli/...
```

确保测试通过，无 vet 警告。

### 3. SKILL.md 规范检查

检查 `skills/md2wechat-lite/SKILL.md` 符合规范：

- [ ] 包含 YAML frontmatter（name + description）
- [ ] name: 小写+连字符，<64 字符
- [ ] description: 第三人称，包含功能+使用场景
- [ ] 使用动名词形式
- [ ] SKILL.md 主体 < 500 行
- [ ] 使用渐进式披露模式
- [ ] 术语一致
- [ ] 引用仅一层深
- [ ] 无时效性信息
- [ ] Unix 路径格式

参考规范：`/root/skill-rules/SKILL-RULE.md`

### 4. 代码与项目一致性检查

- [ ] 代码变更与 SKILL.md 描述一致
- [ ] README.md 与实际功能一致
- [ ] 项目结构文档与实际目录结构一致

### 5. README.md 检查

检查 `README.md` 符合最佳排版和阅读习惯：

- [ ] 项目名 md2wechat-lite 与 CLI 名 md2wx 区分清晰
- [ ] SEO 标题：AI Agent 自动排版公众号 —— Markdown 转微信排版 CLI 工具
- [ ] 前置准备：API Key 和主题预览链接
- [ ] 功能亮点简洁明了
- [ ] AI 创作工作流说明
- [ ] 相关项目链接
- [ ] 打赏和作者信息完整
- [ ] OpenClaw 链接：https://openclaw.ai

### 6. Git Commit 规范

提交信息格式：

```
<type>(<scope>): <subject>

<body>

<footer>
```

**Type**:
- `feat`: 新功能
- `fix`: 修复 bug
- `docs`: 文档变更
- `style`: 代码格式（不影响功能）
- `refactor`: 重构
- `test`: 测试相关
- `chore`: 构建过程或辅助工具变更

**重要**: Commit 中不能有 claude 账号和邮箱信息！

正确示例：
```
feat(api): add newspic draft support

Add newspic draft command for creating image-rich articles.
```

### 7. 更新 CHANGELOG.md

在 `CHANGELOG.md` 顶部添加新版本：

```markdown
## [1.x.x] - 2025-02-27

### Added
- 新增功能说明

### Changed
- 变更说明

### Fixed
- 修复说明

### Removed
- 移除说明
```

### 8. 发版流程

```bash
# 1. 确认所有检查通过
# 2. 提交变更
git add .
git commit -m "chore(release): prepare for v1.x.x"

# 3. 创建标签
git tag -a v1.x.x -m "Release v1.x.x"

# 4. 推送到远程
git push origin main
git push origin v1.x.x
```

## 技术要点

### 配置管理

- 配置文件位置：`~/.md2wx/config.yaml`
- 配置优先级：命令行参数 > 环境变量 > 配置文件 > 默认值
- 环境变量前缀：`MD2WX_`

### 主题系统

- 内置主题（6 种）：default, bytedance, chinese, apple, sports, cyber
- 模板主题（32 种）：{minimal|focus|elegant|bold}-{gold|green|blue|orange|red|navy|gray|sky}

### API 端点

- 默认 API Base URL: `http://111.231.20.31:8080`
- API Key 获取: https://md2wechat.app/api-docs
- 主题预览: https://md2wechat.app/theme-gallery

### 零依赖设计

- 除 Cobra 外无其他依赖
- 手动实现 key=value 配置解析（不依赖 YAML 库）

## 相关链接

- GitHub: https://github.com/geekjourneyx/md2wechat-lite
- API 文档: https://md2wechat.app/api-docs
- 主题画廊: https://md2wechat.app/theme-gallery
- md2wechat-skill: https://github.com/geekjourneyx/md2wechat-skill

## 作者

- **作者**: geekjourneyx
- **X (Twitter)**: https://x.com/seekjourney
- **公众号**: 极客杰尼
