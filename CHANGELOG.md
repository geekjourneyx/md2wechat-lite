# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.0.0] - 2025-02-27

### Added
- Initial release of md2wechat-lite
- Article draft command for creating WeChat article drafts from Markdown
- Newspic draft command for creating Xiaolvshu (image card) drafts
- Batch upload command for uploading images to WeChat CDN
- Theme system with 6 built-in themes and 32 template themes
- Config management commands (set/get/list/path)
- Zero-dependency design (except cobra)
- Support for environment variable configuration
- Claude Code Skill definition
- Installation script for Linux/macOS/Windows

### Features
- 38+ formatting themes
- JSON output format for all commands
- Multiple font size options (small/medium/large)
- Background type options (none/default/grid)
- Global options for API base URL and API key override
- Configuration priority: CLI args > env vars > config file > defaults

[Unreleased]: https://github.com/geekjourneyx/md2wechat-lite/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/geekjourneyx/md2wechat-lite/releases/tag/v1.0.0
