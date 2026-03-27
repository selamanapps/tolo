.PHONY: build install clean test help release lint fmt

BIN_NAME=tolo
BUILD_DIR=build
PREFIX?=/usr/local
DESTDIR=""
VERSION=$(shell cat VERSION)
LDFLAGS=-ldflags="-s -w -X main.version=$(VERSION)"

help:
	@echo "Available targets:"
	@echo "  build      - Build the binary"
	@echo "  install    - Install to $(PREFIX)/bin"
	@echo "  uninstall  - Remove from $(PREFIX)/bin"
	@echo "  clean      - Remove build artifacts"
	@echo "  test       - Run tests"
	@echo "  lint       - Run linter"
	@echo "  fmt        - Format code"
	@echo "  release    - Build release binaries"
	@echo "  help       - Show this help message"

build:
	@echo "Building $(BIN_NAME) v$(VERSION)..."
	@mkdir -p $(BUILD_DIR)
	@go build $(LDFLAGS) -o $(BUILD_DIR)/$(BIN_NAME)
	@echo "Build complete: $(BUILD_DIR)/$(BIN_NAME)"
	@ls -lh $(BUILD_DIR)/$(BIN_NAME)

install: build
	@echo "Installing $(BIN_NAME) v$(VERSION) to $(PREFIX)/bin..."
	@install -m 755 $(BUILD_DIR)/$(BIN_NAME) $(DESTDIR)$(PREFIX)/bin/$(BIN_NAME)
	@echo "Installation complete!"
	@echo ""
	@echo "To enable shell completion, add to your shell config:"
	@echo "  Bash: echo 'source <(tolo --bash-completion)' >> ~/.bashrc"
	@echo "  Zsh:  echo 'source <(tolo --zsh-completion)' >> ~/.zshrc"

uninstall:
	@echo "Removing $(BIN_NAME) from $(PREFIX)/bin..."
	@rm -f $(DESTDIR)$(PREFIX)/bin/$(BIN_NAME)
	@echo "Uninstallation complete!"

clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete!"

test:
	@echo "Running tests..."
	@go test -v ./...

lint:
	@echo "Running linter..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed. Install with: curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b \$$(go env GOPATH)/bin"; \
	fi

fmt:
	@echo "Formatting code..."
	@go fmt ./...
	@goimports -w .

release: clean
	@echo "Building release binaries v$(VERSION)..."
	@mkdir -p $(BUILD_DIR)
	@echo "Building for Linux AMD64..."
	@GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BIN_NAME)-linux-amd64
	@echo "Building for Linux ARM64..."
	@GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BIN_NAME)-linux-arm64
	@echo "Building for macOS AMD64..."
	@GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BIN_NAME)-darwin-amd64
	@echo "Building for macOS ARM64..."
	@GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BIN_NAME)-darwin-arm64
	@echo "Building for Windows AMD64..."
	@GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BIN_NAME)-windows-amd64.exe
	@echo "Building for Windows 386..."
	@GOOS=windows GOARCH=386 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BIN_NAME)-windows-386.exe
	@echo "Release build complete!"
	@ls -lh $(BUILD_DIR)/
