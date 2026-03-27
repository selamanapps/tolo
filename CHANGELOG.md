# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Planned Features
- Export/Import aliases
- Alias categories/tags
- Configuration file support
- Interactive mode with fuzzy search
- Alias history/undo functionality
- Cloud sync across devices
- GUI application

## [1.0.0] - 2026-03-27

### Added
- Initial release of Tolo
- Save aliases with `tolo save` command
- Run saved aliases with `tolo run` command
- Update existing aliases with `tolo update` command
- Delete aliases with `tolo delete` command
- List all aliases with formatted table via `tolo list`
- Show detailed alias information with `tolo show`
- Search aliases via `tolo search`
- Bash and Zsh shell completion
- Command shortcuts (s, r, u, d, ls, l, sh, se, h, v)
- Beautiful terminal UI with colors and icons
- JSON-based storage in `~/.tolo/tolo.db.json`
- Cross-platform support (Linux, macOS, Windows)
- Single binary distribution (~2MB)
- Installation script
- Makefile for build automation
- Comprehensive documentation

### Features
- Lightning-fast execution (Go-based)
- Minimal RAM footprint
- Command parsing with quote handling
- Efficient JSON marshaling
- Error handling with user-friendly messages
- Auto-completion for aliases

[Unreleased]: https://github.com/yourusername/tolo/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/yourusername/tolo/releases/tag/v1.0.0
