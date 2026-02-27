# md2wechat-lite

<div align="center">

**AI Agent 自动排版公众号 —— Markdown 转微信排版 CLI 工具**

打通公众号创作的最后一公里

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GitHub Release](https://img.shields.io/github/v/release/geekjourneyx/md2wechat-lite)](https://github.com/geekjourneyx/md2wechat-lite/releases)

**像发朋友圈一样简单，用 Markdown 写公众号文章**

md2wx CLI · [38+ 精美主题](https://md2wechat.app/theme-gallery) · 支持 [Claude Code](https://claude.ai/code) / [OpenClaw](https://openclaw.ai)

</div>

---

## 前置准备

使用前请先获取 API Key：

- [获取 API Key](https://md2wechat.app/api-docs)
- [主题预览](https://md2wechat.app/theme-gallery)

---

## 快速开始

### 安装

```bash
curl -fsSL https://raw.githubusercontent.com/geekjourneyx/md2wechat-lite/main/cli/scripts/install.sh | sh
```

### 配置

```bash
md2wx config set wechat-appid "wx123..."
md2wx config set wechat-appsecret "your_secret"
md2wx config set api-key "wme_your_api_key"
```

### 第一个草稿

```bash
md2wx article-draft --markdown "# 欢迎\n\n这是我的第一篇文章！"
```

---

## 功能亮点

### 📝 图文草稿

Markdown 转 WeChat 排版，一键生成草稿

```bash
md2wx article-draft --file article.md --theme bytedance
```

### 🖼️ 小绿书草稿

创建图片文章，支持多图上传

```bash
md2wx newspic-draft --title "标题" --content "内容" --images img1.jpg,img2.png
```

### 📦 批量上传

上传图片到微信素材库，获取永久 URL

```bash
md2wx batch-upload --images *.jpg
```

### 🎨 38+ 主题

- **内置 6 种**：default, bytedance, chinese, apple, sports, cyber
- **模板 32 种**：{minimal|focus|elegant|bold} × {gold|green|blue|orange|red|navy|gray|sky}

```bash
md2wx themes list --verbose
```

---

## AI 创作工作流

### Claude Code

配合 [md2wechat-skill](https://github.com/geekjourneyx/md2wechat-skill) 使用：

```
1. 在 Claude Code 中激活 Skill
2. "帮我写一篇关于 AI 的文章"
3. "发布到公众号草稿箱"
```

### OpenClaw

支持 OpenClaw 自动化工作流

---

## 配置说明

配置文件：`~/.md2wx/config.yaml`

**配置优先级**：命令行参数 > 环境变量 > 配置文件 > 默认值

```bash
# 查看配置
md2wx config list

# 设置默认主题
md2wx config set default-theme "bytedance"

# 设置字体大小
md2wx config set font-size "large"
```

---

## 常见问题

<details>
<summary>安装后提示 command not found？</summary>

将二进制文件目录添加到 PATH，或打开新终端窗口
</details>

<details>
<summary>macOS 无法验证开发者？</summary>

```bash
chmod +x md2wx
xattr -d com.apple.quarantine md2wx
```
</details>

<details>
<summary>草稿创建成功但后台找不到？</summary>

检查 AppID/AppSecret 是否正确，确认登录了对应公众号账号
</details>

---

## 开发

```bash
git clone https://github.com/geekjourneyx/md2wechat-lite.git
cd md2wechat-lite
go build -o md2wx ./cli
```

---

## 相关项目

- [md2wechat-skill](https://github.com/geekjourneyx/md2wechat-skill) - 用 Markdown 写公众号文章，像发朋友圈一样简单

---

## License

[MIT License](LICENSE)

---

## 💰 打赏 Buy Me A Coffee

如果该项目帮助了您，请作者喝杯咖啡吧 ☕️

### WeChat

<img src="https://raw.githubusercontent.com/geekjourneyx/awesome-developer-go-sail/main/docs/assets/wechat-reward-code.jpg" alt="微信打赏码" width="200" />

---

## 🧑‍💻 作者

- **作者**：[geekjourneyx](https://geekjourney.dev)
- **X (Twitter)**：https://x.com/seekjourney
- **公众号**：极客杰尼

关注公众号，获取更多 AI 编程、AI 工具与 AI 出海建站的实战分享：

<p align="center">
<img src="https://raw.githubusercontent.com/geekjourneyx/awesome-developer-go-sail/main/docs/assets/qrcode.jpg" alt="公众号：极客杰尼" width="180" />
</p>

---

<div align="center">

**让公众号写作更简单** ⭐

Made with ❤️ by [geekjourneyx](https://geekjourney.dev)

</div>
