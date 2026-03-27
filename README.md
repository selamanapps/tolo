<div align="center">

# 🚀 Tolo

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-green.svg)
![Platform](https://img.shields.io/badge/platform-linux%20%7C%20macos%20%7C%20windows-lightgrey)
![Size](https://img.shields.io/badge/Size-~2MB-blue)

### Lightning-fast command alias manager for the modern terminal

[Install](#-installation) • [Features](#-features) • [Usage](#-usage) • [Contributing](#-contributing)

</div>

---

## ✨ Features

- 🚀 **Blazing Fast** - Written in Go, executes in milliseconds
- 💾 **Lightweight** - Only ~2MB binary, minimal RAM footprint
- 🎨 **Beautiful UI** - Colorful, icon-rich terminal output
- 🔍 **Search** - Find aliases instantly with fuzzy search
- 🔄 **Update** - Modify aliases on the fly
- 🗑️ **Delete** - Remove unwanted aliases
- 📋 **List** - View all saved aliases in a formatted table
- ⚡ **Shortcuts** - Use short commands like `ls`, `rm`, `s`, `r`
- 🔧 **Shell Completion** - Bash and Zsh auto-completion
- 📦 **Single Binary** - No dependencies, just copy and run
- 🌐 **Cross-platform** - Linux, macOS, and Windows support

## 🎯 Why Tolo?

Tired of typing long SSH commands, gcloud commands, or complex terminal commands? Tolo saves them as simple aliases you can run with a single command.

**Before:**
```bash
ssh user@192.168.1.10 -p 2222 -i ~/.ssh/mykey.pem
```

**After:**
```bash
tolo run myserver
# or shorter
tolo r myserver
```

## 📦 Installation

### Quick Install (Linux/macOS)

```bash
curl -fsSL https://raw.githubusercontent.com/selamanapps/tolo/main/install.sh | bash
```

### Manual Install

#### Download Binary

Visit the [Releases](https://github.com/selamanapps/tolo/releases) page and download the binary for your platform.

```bash
# Linux
wget https://github.com/selamanapps/tolo/releases/download/v1.0.0/tolo-linux-amd64
sudo cp tolo-linux-amd64 /usr/local/bin/tolo
sudo chmod +x /usr/local/bin/tolo

# macOS
curl -L https://github.com/selamanapps/tolo/releases/download/v1.0.0/tolo-darwin-amd64 -o tolo
sudo cp tolo /usr/local/bin/
sudo chmod +x /usr/local/bin/tolo
```

#### Build from Source

```bash
git clone https://github.com/selamanapps/tolo.git
cd tolo
go build -ldflags="-s -w" -o tolo
sudo cp tolo /usr/local/bin/
```

#### Shell Completion

Enable auto-completion for your shell:

**Bash:**
```bash
echo 'source <(tolo --bash-completion)' >> ~/.bashrc
source ~/.bashrc
```

**Zsh:**
```bash
echo 'source <(tolo --zsh-completion)' >> ~/.zshrc
source ~/.zshrc
```

## 🎮 Usage

### Basic Commands

```bash
# Save a new alias
tolo save server1:ssh user@192.168.1.10

# Run a saved alias
tolo run server1

# List all aliases
tolo list

# Delete an alias
tolo delete server1

# Search aliases
tolo search ssh
```

### Shortcuts (Power User Commands)

```bash
# Short aliases for all commands
tolo s   # save
tolo r   # run
tolo u   # update
tolo d   # delete (also: del, rm)
tolo ls  # list (also: l)
tolo sh  # show (also: info)
tolo se  # search (also: find)
```

### Examples

#### SSH Connections

```bash
# Save SSH connection
tolo save mypc:ssh amancca@192.168.0.100

# Use it
tolo r mypc
```

#### Cloud Commands

```bash
# Save gcloud command
tolo save ai-server:gcloud compute ssh --zone us-central1-c ai-agent --project my-project

# Execute it
tolo r ai-server
```

#### Docker Commands

```bash
# Save complex docker command
tolo save dev:docker-compose up -d --build

# Run it
tolo r dev
```

#### Update Existing Alias

```bash
# Update to change connection details
tolo u mypc:ssh root@192.168.0.100
```

#### Show Alias Details

```bash
tolo show mypc
```

#### Search Aliases

```bash
# Find all SSH aliases
tolo se ssh

# Find all docker aliases
tolo find docker
```

## 📸 Screenshots

<div align="center">

**List Command**
```
╔════════════════════════════════════════════════════════════════╗
║                    📋 Saved Aliases                            ║
╚════════════════════════════════════════════════════════════════╝

  1  mypc       →  ssh amancca@192.168.0.100
  2  ai-server  →  gcloud compute ssh ai-agent --project my-journey-app-482201
  3  dev        →  docker-compose up -d --build

──────────────────────────────────────────────────────────────────
  Total: 3
```

**Save Command**
```
💾  Alias saved successfully

Alias:   mypc
Command: ssh amancca@192.168.0.100
──────────────────────────────────────────────────────────────────
```

</div>

## 🛠️ Development

### Build

```bash
# Build for current platform
make build

# Build for all platforms
make build-all

# Install to system
make install
```

### Project Structure

```
tolo/
├── cmd/              # Command handlers
├── storage/          # JSON file operations
├── executor/         # Command execution
├── pretty/           # Terminal formatting
├── completion/       # Shell completions
├── main.go          # Entry point
├── Makefile         # Build automation
└── install.sh       # Installation script
```

## 🤝 Contributing

Contributions are welcome! Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 Roadmap

- [ ] Export/Import aliases
- [ ] Alias categories/tags
- [ ] Configuration file support
- [ ] Interactive mode
- [ ] Alias history/undo
- [ ] Sync across devices
- [ ] GUI application

## 🐧 Requirements

- Go 1.21 or higher (for building from source)
- Linux, macOS, or Windows

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Inspired by tools like [alias](https://wiki.archlinux.org/title/Alias) and [gnu stow](https://www.gnu.org/software/stow/)
- Built with [Go](https://golang.org/)
- Icons and colors for better UX

## 📞 Support

- 📖 [Documentation](https://github.com/selamanapps/tolo/wiki)
- 🐛 [Issue Tracker](https://github.com/selamanapps/tolo/issues)
- 💬 [Discussions](https://github.com/selamanapps/tolo/discussions)

---

<div align="center">

Made with ❤️ at Zemenawi Lab

[⬆ Back to top](#-tolo)

</div>
