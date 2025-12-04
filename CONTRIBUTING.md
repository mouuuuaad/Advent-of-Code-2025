# Contributing to Advent of Code 2025 Solutions

Thank you for your interest! While this is primarily a personal project for solving Advent of Code challenges, contributions are welcome.

## 🎯 Philosophy

This repository demonstrates:
- Clean, readable code across multiple languages
- Consistent project structure
- Comprehensive testing
- Clear documentation

## 📝 Adding Solutions

### File Structure

Each solution should follow this structure:

**Go** (`go/dayXX/`):
```
main.go       # Solution implementation
main_test.go  # Unit tests
README.md     # Problem description and approach
```

**Python** (`python/dayXX/`):
```
solution.py      # Solution implementation
test_solution.py # Unit tests
README.md        # Problem description and approach
```

### Code Style

#### Go
- Follow standard Go formatting (`gofmt`)
- Use meaningful variable names
- Add comments for complex logic
- Write table-driven tests
- Avoid panics in production code (use error returns)

#### Python
- Follow PEP 8 style guide
- Use type hints where helpful
- Write docstrings for functions
- Use pytest for testing
- Keep functions focused and single-purpose

### Testing Requirements

All solutions must include:
1. **Example tests**: Test with provided example inputs
2. **Edge cases**: Test boundary conditions
3. **Documentation**: Explain the approach in README

Run tests before submitting:
```bash
# Go
cd go && go test ./...

# Python
cd python && pytest
```

## 🔧 Development Workflow

1. **Create solution files** using the existing templates
2. **Implement the solution** with clear logic
3. **Add tests** with example inputs
4. **Document your approach** in the day's README
5. **Run tests** to verify correctness
6. **Update main README** progress table

## 📚 Commit Conventions

Use clear, descriptive commit messages:

```
Add solution for day 5 part 1
Optimize day 3 solution
Fix edge case in day 7
Update README with day 10 progress
```

## 🚫 What Not to Include

- **Input files** (`input.txt`) - These are gitignored per AoC policy
- **Compiled binaries** - Already in `.gitignore`
- **Virtual environments** - Already in `.gitignore`
- **IDE-specific files** - Already in `.gitignore`

## ✅ Quality Checklist

Before submitting:

- [ ] Code follows language-specific style guidelines
- [ ] Tests pass (`make test-all`)
- [ ] README updated with approach and complexity
- [ ] No input files committed
- [ ] Code is well-commented
- [ ] CI passes on GitHub Actions

## 🤝 Questions?

Feel free to open an issue for:
- Clarification on project structure
- Discussion of alternative approaches
- Bug reports
- Feature suggestions

## 📜 License

By contributing, you agree that your contributions will be licensed under the MIT License.

---

Happy coding! 🎄⭐
