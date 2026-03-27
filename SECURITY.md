# Security Policy

## Supported Versions

Only the latest version of Tolo receives security updates and bug fixes. Users are strongly encouraged to keep Tolo updated to the latest version.

## Reporting a Vulnerability

If you discover a security vulnerability in Tolo, please report it responsibly.

### How to Report

1. **Do not** create a public issue
2. Send an email to: [security@selamanapps.com](mailto:security@selamanapps.com)
3. Include as much detail as possible:
   - Description of the vulnerability
   - Steps to reproduce
   - Potential impact
   - Suggested fix (if any)

### What Happens Next?

- You will receive an acknowledgment within 48 hours
- We will investigate the vulnerability
- We will work with you to develop a fix
- Once fixed, we will coordinate the disclosure

## Security Best Practices

### File Permissions

Tolo stores aliases in `~/.tolo/tolo.db.json`. Ensure this file has appropriate permissions:

```bash
chmod 600 ~/.tolo/tolo.db.json
```

### Sensitive Information

- Avoid storing passwords or API keys directly in aliases
- Use environment variables for sensitive data
- Be careful with commands that contain credentials

### Command Execution

Tolo executes commands exactly as saved. Always verify aliases before running:

```bash
tolo show alias-name
```

## Dependency Security

Tolo is built with pure Go and minimal dependencies. We regularly update dependencies to address security issues. The project uses GitHub Dependabot for automated dependency updates.

## Security Features

- **No remote network calls** - Tolo runs entirely locally
- **File-based storage** - No database servers
- **Simple JSON format** - Easy to audit
- **No external dependencies** - Minimal attack surface
