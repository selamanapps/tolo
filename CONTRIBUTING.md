# Contributing to Tolo

Thank you for considering contributing to Tolo! This document provides guidelines and instructions for contributing to the project.

## 🤝 How to Contribute

### Reporting Bugs

Before creating bug reports, please check the existing issues as you might find that the problem has already been reported.

When creating bug reports, please include:

- **Clear title and description**
- **Steps to reproduce** the issue
- **Expected behavior** vs **actual behavior**
- **Environment details**: OS, Go version, Tolo version
- **Screenshots** if applicable
- **Error messages** or logs

### Suggesting Enhancements

Enhancement suggestions are tracked as GitHub issues. When suggesting:

- Use a clear and descriptive title
- Provide detailed explanation of the enhancement
- Explain why it would be useful
- Provide examples or use cases

## 🛠️ Development Setup

1. **Fork the repository**
   ```bash
   # Fork the repo on GitHub, then clone your fork
   git clone https://github.com/YOUR_USERNAME/tolo.git
   cd tolo
   ```

2. **Add upstream remote**
   ```bash
   git remote add upstream https://github.com/selamanapps/tolo.git
   ```

3. **Create a feature branch**
   ```bash
   git checkout -b feature/your-feature-name
   ```

4. **Build and test**
   ```bash
   make build
   make test
   ```

## 📝 Coding Standards

### Go Conventions

- Follow standard Go conventions defined in [Effective Go](https://golang.org/doc/effective_go.html)
- Run `gofmt` on your code before committing
- Use `golint` to check for potential issues
- Write clear, concise comments for exported functions

### Code Style

- Keep functions small and focused
- Use meaningful variable and function names
- Add error handling where appropriate
- Write tests for new features
- Keep the binary size minimal

### Project Structure

- Place command handlers in `cmd/`
- Storage operations go in `storage/`
- Terminal formatting in `pretty/`
- Follow the existing package structure

## 🧪 Testing

```bash
# Run all tests
make test

# Run with coverage
go test -cover ./...

# Run specific package tests
go test ./cmd
```

## 📦 Building

```bash
# Build for current platform
make build

# Build for all platforms
make build-all

# Run locally
./build/tolo help
```

## 🚀 Pull Request Process

1. **Update documentation** if needed
2. **Add tests** for new features or bug fixes
3. **Run tests** and ensure they pass
4. **Update README.md** if you add features
5. **Push to your fork**
6. **Create a Pull Request**

### Pull Request Guidelines

- Use a clear title describing the change
- Reference related issues in the description
- Include screenshots for UI changes
- Keep changes focused and minimal
- Respond to review comments promptly

## 📋 Commit Messages

Follow [Conventional Commits](https://www.conventionalcommits.org/):

- `feat:` New feature
- `fix:` Bug fix
- `docs:` Documentation changes
- `style:` Code style changes (formatting, etc.)
- `refactor:` Code refactoring
- `perf:` Performance improvements
- `test:` Adding or updating tests
- `chore:` Maintenance tasks

Examples:
```
feat: add export command for aliases
fix: handle special characters in commands
docs: update installation instructions
```

## 🌍 Localization

Tolo is currently in English. If you'd like to add translations:

1. Create a new issue to discuss
2. Follow i18n best practices
3. Test all translated text
4. Submit a PR with translation files

## 📚 Documentation

- Update README.md for user-facing changes
- Add comments to complex code
- Update inline documentation
- Consider adding wiki pages for complex features

## 🎯 Feature Development

Before implementing major features:

1. Open an issue to discuss
2. Get feedback from maintainers
3. Design the feature
4. Create a proposal PR (optional)
5. Implement and test

## ⚠️ Breaking Changes

Avoid breaking changes if possible. If necessary:

- Clearly document the change
- Update version numbers
- Provide migration guide
- Discuss with maintainers first

## 🤖 Automation

The project uses:

- **Make** for build automation
- **GitHub Actions** for CI/CD
- **gofmt** for code formatting
- **golint** for linting

## 📄 License

By contributing, you agree that your contributions will be licensed under the MIT License.

## 💬 Getting Help

- Check [Documentation](README.md)
- Search [Issues](../../issues)
- Start a [Discussion](../../discussions)
- Ask in PR comments

## 🎉 Recognition

Contributors will be recognized in:

- Contributors section
- Release notes
- Project documentation

Thank you for contributing to Tolo! 🚀
