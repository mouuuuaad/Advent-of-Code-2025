# 🎄 Advent of Code 2025 - Multi-Language Solutions

My solutions to [Advent of Code 2025](https://adventofcode.com/2025) implemented in multiple programming languages!

## 📁 Repository Structure

This repository showcases solutions in different programming languages, with 4-5 days per language:

```
├── go/            # Days 1-4   (Go) ✅
├── python/        # Days 5-9   (Python) - Structure ready
├── typescript/    # Days 10-14 (TypeScript/Node.js) - Structure ready
└── javascript/    # Days 15-19 (JavaScript) - Structure ready
```

Each language directory contains:
- Individual `dayXX/` folders with solution files and tests
- `README.md` files explaining each problem
- Shared utilities for common operations

## 🚀 Quick Start

### Go Solutions (Days 1-4)

**Prerequisites**: Go 1.21 or higher

```bash
# Run all tests
cd go && go test ./...

# Run specific day
cd go/day01 && go run main.go

# Using Makefile
make go-test
make go-run DAY=1
```

### Python Solutions (Days 5-9)

**Prerequisites**: Python 3.10 or higher

**Status**: Directory structure ready, awaiting problem releases.

```bash
# Setup virtual environment (when ready to start)
make python-setup

# Each day will have:
# - solution.py (main solution)
# - test_solution.py (unit tests)
# - README.md (problem description)
```

### TypeScript Solutions (Days 10-14)

**Prerequisites**: Node.js 18+ and npm

**Status**: Directory structure ready, awaiting problem releases.

```bash
# Each day will have:
# - solution.ts (main solution)
# - test.ts (unit tests)
# - README.md (problem description)
```

## 📊 Progress

| Day | Language   | Part 1 | Part 2 | Solution |
|-----|------------|--------|--------|----------|
| 01  | Go         | ⭐     | ⭐     | [go/day01](go/day01) |
| 02  | Go         | ⭐     | ⭐     | [go/day02](go/day02) |
| 03  | Go         | ⭐     | ⭐     | [go/day03](go/day03) |
| 04  | Go         | ⭐     | ⭐     | [go/day04](go/day04) |
| 05  | Python     | 📁     | 📁     | [python/day05](python/day05) |
| 06  | Python     | 📁     | 📁     | [python/day06](python/day06) |
| 07  | Python     | 📁     | 📁     | [python/day07](python/day07) |
| 08  | Python     | 📁     | 📁     | [python/day08](python/day08) |
| 09  | Python     | 📁     | 📁     | [python/day09](python/day09) |
| 10  | TypeScript | 📁     | 📁     | [typescript/day10](typescript/day10) |
| 11  | TypeScript | 📁     | 📁     | [typescript/day11](typescript/day11) |
| 12  | TypeScript | 📁     | 📁     | [typescript/day12](typescript/day12) |
| 13  | TypeScript | 📁     | 📁     | [typescript/day13](typescript/day13) |
| 14  | TypeScript | 📁     | 📁     | [typescript/day14](typescript/day14) |
| ... | ...        | ...    | ...    | ... |

**Legend**: ⭐ Completed | 📁 Structure ready | ⬜ Not started

## 🛠️ Development

### Run All Tests

```bash
make test-all
```

### Clean Build Artifacts

```bash
make clean
```

### View All Available Commands

```bash
make help
```

## 📝 Structure Details

Each language directory follows a consistent structure:

**Go** (`go/dayXX/`):
```
main.go       # Solution implementation
main_test.go  # Unit tests
README.md     # Problem description
```

**Python** (`python/dayXX/`):
```
solution.py      # Solution implementation (when started)
test_solution.py # Unit tests (when started)
README.md        # Problem description
```

**TypeScript** (`typescript/dayXX/`):
```
solution.ts  # Solution implementation (when started)
test.ts      # Unit tests (when started)
README.md    # Problem description
```

All languages include shared utility modules in their respective directories.

## 🧪 Testing

All solutions include:
- Unit tests with example inputs
- Automated CI via GitHub Actions
- Test coverage reporting (Python)
- Race condition detection (Go)

## 📜 License

MIT License - see [LICENSE](LICENSE) file for details

## 🔗 Links

- [Advent of Code 2025](https://adventofcode.com/2025)
- [About Advent of Code](https://adventofcode.com/about)

---

**Note**: Input files are not included in this repository per Advent of Code's [terms of use](https://adventofcode.com/about#faq_copying). Please download your own inputs from the [Advent of Code website](https://adventofcode.com).
