# Contributing to Motor Finance Complaint Generator

Thank you for your interest in contributing! This project welcomes contributions from the community.

## Getting Started

1. **Fork the repository** on GitHub
2. **Clone your fork** locally:
   ```bash
   git clone https://github.com/YOUR_USERNAME/synthetic-emails.git
   cd synthetic-emails
   ```
3. **Install dependencies**:
   ```bash
   go mod tidy
   ```
4. **Create a feature branch**:
   ```bash
   git checkout -b feature/your-feature-name
   ```

## 🛠️ Development Setup

### Prerequisites

- Go 1.21 or higher
- Git

### Running Tests

```bash
make test
```

### Running the Application

```bash
make demo
```

### Building

```bash
make build
```

## 📝 Making Changes

### Code Style

- Follow standard Go conventions
- Use `gofmt` to format your code
- Add comments for exported functions
- Write tests for new functionality

### Commit Messages

Use clear, descriptive commit messages:

```
feat: add support for custom email templates
fix: resolve PDF password generation issue
docs: update README with new features
```

### Pull Request Process

1. **Update documentation** if needed
2. **Add tests** for new functionality
3. **Ensure all tests pass**: `make test`
4. **Update README.md** if adding new features
5. **Create a Pull Request** with a clear description

### Pull Request Template

```markdown
## Description

Brief description of changes

## Type of Change

- [ ] Bug fix
- [ ] New feature
- [ ] Documentation update
- [ ] Performance improvement

## Testing

- [ ] Tests pass locally
- [ ] Added tests for new functionality
- [ ] Tested manually

## Checklist

- [ ] Code follows project style guidelines
- [ ] Self-review completed
- [ ] Documentation updated
```

## 🐛 Reporting Issues

When reporting issues, please include:

- Go version (`go version`)
- Operating system
- Steps to reproduce
- Expected vs actual behavior
- Any error messages

## 💡 Feature Requests

We welcome feature requests! Please:

- Check existing issues first
- Describe the use case
- Explain why it would be valuable
- Consider submitting a PR if you can implement it

## 📚 Documentation

Help improve documentation by:

- Fixing typos or unclear instructions
- Adding examples
- Improving code comments
- Updating README.md

## 🏆 Recognition

Contributors will be recognized in:

- GitHub contributors list
- Release notes (for significant contributions)
- README acknowledgments

## ❓ Questions?

Feel free to:

- Open an issue for questions
- Start a discussion on GitHub
- Reach out to maintainers

Thank you for contributing! 🙏
