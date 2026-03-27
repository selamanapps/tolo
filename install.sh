#!/bin/bash
# Tolo Installation Script

set -e

VERSION="1.0.0"
BINARY_NAME="tolo"
INSTALL_DIR="/usr/local/bin"
REPO_URL="https://github.com/selamanapps/tolo"

detect_os() {
    OS="$(uname -s)"
    ARCH="$(uname -m)"

    case "$OS" in
        Linux*)  OS="linux" ;;
        Darwin*) OS="darwin" ;;
        MINGW*|MSYS*|CYGWIN*) OS="windows" ;;
        *) echo "Unsupported OS: $OS"; exit 1 ;;
    esac

    case "$ARCH" in
        x86_64|amd64) ARCH="amd64" ;;
        i386|i686) ARCH="386" ;;
        arm64|aarch64) ARCH="arm64" ;;
        arm*) ARCH="arm" ;;
        *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
    esac

    echo "$OS-$ARCH"
}

install_from_source() {
    echo "Installing from source..."

    if ! command -v go &> /dev/null; then
        echo "Error: Go is not installed. Please install Go first."
        exit 1
    fi

    TEMP_DIR=$(mktemp -d)
    cd "$TEMP_DIR"

    echo "Downloading source..."
    git clone "$REPO_URL.git" .
    
    echo "Building $BINARY_NAME..."
    go build -ldflags="-s -w -X main.version=$VERSION" -o "$BINARY_NAME"
    
    echo "Installing to $INSTALL_DIR..."
    sudo cp "$BINARY_NAME" "$INSTALL_DIR/"
    sudo chmod +x "$INSTALL_DIR/$BINARY_NAME"
    
    cd - > /dev/null
    rm -rf "$TEMP_DIR"
}

install_prebuilt() {
    PLATFORM=$(detect_os)
    DOWNLOAD_URL="$REPO_URL/releases/download/v$VERSION/$BINARY_NAME-$PLATFORM"
    
    echo "Installing pre-built binary for $PLATFORM..."
    echo "Downloading from: $DOWNLOAD_URL"

    TEMP_DIR=$(mktemp -d)
    cd "$TEMP_DIR"
    
    if command -v curl &> /dev/null; then
        curl -L -o "$BINARY_NAME" "$DOWNLOAD_URL"
    elif command -v wget &> /dev/null; then
        wget -O "$BINARY_NAME" "$DOWNLOAD_URL"
    else
        echo "Error: Neither curl nor wget is installed."
        exit 1
    fi
    
    chmod +x "$BINARY_NAME"
    
    echo "Installing to $INSTALL_DIR..."
    sudo cp "$BINARY_NAME" "$INSTALL_DIR/"
    
    cd - > /dev/null
    rm -rf "$TEMP_DIR"
}

main() {
    echo "╔══════════════════════════════════════════════════════╗"
    echo "║         Tolo Installation Script v$VERSION           ║"
    echo "╚══════════════════════════════════════════════════════╝"
    echo ""
    
    if [ -e "$INSTALL_DIR/$BINARY_NAME" ]; then
        echo "⚠ Warning: $BINARY_NAME is already installed at $INSTALL_DIR/$BINARY_NAME"
        read -p "Do you want to overwrite it? (y/N): " -n 1 -r
        echo ""
        if [[ ! $REPLY =~ ^[Yy]$ ]]; then
            echo "Installation cancelled."
            exit 0
        fi
    fi

    echo "Select installation method:"
    echo "  1) Install pre-built binary (faster)"
    echo "  2) Build from source (requires Go)"
    read -p "Your choice (default: 1): " choice
    
    case "${choice:-1}" in
        1)
            if curl -sSf -I "$REPO_URL/releases/download/v$VERSION/$BINARY_NAME-$(detect_os)" > /dev/null 2>&1; then
                install_prebuilt
            else
                echo "⚠ Pre-built binary not found. Building from source..."
                install_from_source
            fi
            ;;
        2)
            install_from_source
            ;;
        *)
            echo "Invalid choice. Exiting."
            exit 1
            ;;
    esac
    
    echo ""
    echo "✓ Installation completed successfully!"
    echo ""
    echo "To enable shell completion, add to your shell config:"
    echo "  Bash: echo 'source <($BINARY_NAME --bash-completion)' >> ~/.bashrc"
    echo "  Zsh:  echo 'source <($BINARY_NAME --zsh-completion)' >> ~/.zshrc"
    echo ""
    echo "Run '$BINARY_NAME help' to get started."
}

main "$@"
