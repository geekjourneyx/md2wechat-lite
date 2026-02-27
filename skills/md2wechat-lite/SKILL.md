---
name: md2wechat-lite
description: Manages md2wx CLI for converting Markdown to WeChat Official Account drafts. Use when working with WeChat article publishing, Markdown to HTML conversion, or WeChat media upload.
---

# md2wx - Markdown to WeChat

CLI tool for converting Markdown to WeChat Official Account formatted drafts.

## Quick start

**Install**:
```bash
curl -fsSL https://raw.githubusercontent.com/geekjourneyx/md2wechat-lite/main/cli/scripts/install.sh | sh
```

**Configure credentials**:
```bash
md2wx config set wechat-appid "wx123..."
md2wx config set wechat-appsecret "your_secret"
md2wx config set api-key "wme_your_key"
```

## Commands

| Command | Purpose |
|---------|---------|
| `article-draft` | Create article draft from Markdown |
| `newspic-draft` | Create Xiaolvshu (image card) draft |
| `batch-upload` | Upload images to WeChat CDN |
| `themes list` | List available themes |
| `config` | Manage settings (set/get/list/path) |

## Article draft

Convert Markdown to WeChat article:

```bash
md2wx article-draft --file article.md --theme bytedance
```

Or pipe content:
```bash
cat article.md | md2wx article-draft --theme elegant-red
```

## Newspic draft

Create image-rich card drafts:

```bash
md2wx newspic-draft --title "标题" --content "内容" --images img1.jpg,img2.png
```

## Batch upload

Upload images and get WeChat CDN URLs:

```bash
md2wx batch-upload --images *.jpg
```

## Themes

**Built-in** (6): default, bytedance, chinese, apple, sports, cyber

**Template** (32): `{minimal|focus|elegant|bold}-{gold|green|blue|orange|red|navy|gray|sky}`

List/search themes:
```bash
md2wx themes list [--verbose] [--search query]
```

For theme descriptions: See `cli/pkg/themes/list.go`

## Configuration

Config file: `~/.md2wx/config.yaml`

**Priority**: Command args > Environment vars > Config file > Defaults

**Environment variables**:
- `MD2WX_WECHAT_APPID`
- `MD2WX_WECHAT_APPSECRET`
- `MD2WX_API_KEY`
- `MD2WX_API_BASE_URL`
- `MD2WX_DEFAULT_THEME`

## Project structure

```
cli/
├── main.go              # Root command
├── article-draft.go     # Article draft
├── newspic-draft.go     # Xiaolvshu draft
├── batch-upload.go      # Image upload
├── config.go            # Config management
├── themes.go            # Theme list command
└── pkg/
    ├── api/client.go    # HTTP API client
    ├── config/          # Config file I/O
    ├── themes/          # Theme definitions
    └── output/          # JSON formatter
```

## Output format

All commands output JSON:
```json
{
  "success": true,
  "data": { "media_id": "...", "url": "..." }
}
```

## Implementation details

- **Zero dependencies** (except cobra): Manual key=value config parsing
- **Go 1.24+** required
- **Single binary** distribution

See source files for:
- API client: `cli/pkg/api/client.go`
- Config handling: `cli/pkg/config/config.go`
- Theme list: `cli/pkg/themes/list.go`
