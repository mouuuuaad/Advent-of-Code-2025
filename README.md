# 🎄 Advent of Code 2025 - Multi-Language Solutions

My solutions to [Advent of Code 2025](https://adventofcode.com/2025) implemented in multiple programming languages!

## 📁 Repository Structure

This repository showcases solutions in different programming languages, with 4-5 days per language:

```
├── go/            # Days 1-4   (Go) ✅
├── c/             # Days 5-9   (C) - Structure ready
├── cpp/           # Days 10-14 (C++) - Structure ready
└── odin/          # Days 15-19 (Odin) - Structure ready
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
make go-run DAY=1
```

### C Solutions (Days 5-9)

**Prerequisites**: GCC or Clang

**Status**: Directory structure ready, awaiting problem releases.

```bash
# Compile and run
cd c/day05
gcc -o solution main.c ../utils/input.c
./solution
```

### C++ Solutions (Days 10-14)

**Prerequisites**: GCC/Clang with C++20 support

**Status**: Directory structure ready, awaiting problem releases.

```bash
# Compile and run
cd cpp/day10
g++ -std=c++20 -o solution main.cpp -I../utils
./solution
```

### Odin Solutions (Days 15-19)

**Prerequisites**: Odin compiler

**Status**: Directory structure ready, awaiting problem releases.

```bash
# Run
cd odin/day15
odin run .
```

## 📊 Progress

| Day | Language | Part 1 | Part 2 | Solution |
|-----|----------|--------|--------|----------|
| 01  | Go       | ⭐     | ⭐     | [go/day01](go/day01) |
| 02  | Go       | ⭐     | ⭐     | [go/day02](go/day02) |
| 03  | Go       | ⭐     | ⭐     | [go/day03](go/day03) |
| 04  | Go       | ⭐     | ⭐     | [go/day04](go/day04) |
| 05  | C        | 📁     | 📁     | [c/day05](c/day05) |
| 06  | C        | 📁     | 📁     | [c/day06](c/day06) |
| 07  | C        | 📁     | 📁     | [c/day07](c/day07) |
| 08  | C        | 📁     | 📁     | [c/day08](c/day08) |
| 09  | C        | 📁     | 📁     | [c/day09](c/day09) |
| 10  | C++      | 📁     | 📁     | [cpp/day10](cpp/day10) |
| 11  | C++      | 📁     | 📁     | [cpp/day11](cpp/day11) |
| 12  | C++      | 📁     | 📁     | [cpp/day12](cpp/day12) |
| 13  | C++      | 📁     | 📁     | [cpp/day13](cpp/day13) |
| 14  | C++      | 📁     | 📁     | [cpp/day14](cpp/day14) |
| 15  | Odin     | 📁     | 📁     | [odin/day15](odin/day15) |
| 16  | Odin     | 📁     | 📁     | [odin/day16](odin/day16) |
| 17  | Odin     | 📁     | 📁     | [odin/day17](odin/day17) |
| 18  | Odin     | 📁     | 📁     | [odin/day18](odin/day18) |
| 19  | Odin     | 📁     | 📁     | [odin/day19](odin/day19) |
| ... | ...      | ...    | ...    | ... |

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

**C** (`c/dayXX/`):
```
main.c        # Solution implementation
README.md     # Problem description
```

**C++** (`cpp/dayXX/`):
```
main.cpp      # Solution implementation
README.md     # Problem description
```

**Odin** (`odin/dayXX/`):
```
main.odin     # Solution implementation
README.md     # Problem description
```

All languages include shared utility modules in their respective directories.

## 📜 License

MIT License - see [LICENSE](LICENSE) file for details

## 🔗 Links

- [Advent of Code 2025](https://adventofcode.com/2025)
- [About Advent of Code](https://adventofcode.com/about)

---

**Note**: Input files are not included in this repository per Advent of Code's [terms of use](https://adventofcode.com/about#faq_copying). Please download your own inputs from the [Advent of Code website](https://adventofcode.com).
