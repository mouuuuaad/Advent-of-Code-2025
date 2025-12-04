.PHONY: help go-test go-run go-bench python-test python-run python-setup test-all clean

# Default target
help:
	@echo "Advent of Code 2025 - Multi-Language Solutions"
	@echo ""
	@echo "Available targets:"
	@echo "  go-test          - Run all Go tests"
	@echo "  go-run DAY=X     - Run specific Go day (1-4)"
	@echo "  go-bench         - Run Go benchmarks"
	@echo "  python-test      - Run all Python tests"
	@echo "  python-run DAY=X - Run specific Python day (5-9)"
	@echo "  python-setup     - Setup Python virtual environment"
	@echo "  test-all         - Run all tests (Go + Python)"
	@echo "  clean            - Clean build artifacts"

# Go targets
go-test:
	@echo "Running Go tests..."
	cd go && go test ./...

go-run:
	@if [ -z "$(DAY)" ]; then echo "Usage: make go-run DAY=1"; exit 1; fi
	@printf -v day "%02d" $(DAY)
	@echo "Running Go day $(DAY)..."
	cd go/day$${day} && go run main.go

go-bench:
	@echo "Running Go benchmarks..."
	cd go && go test -bench=. ./...

# Python targets
python-test:
	@echo "Running Python tests..."
	cd python && python -m pytest

python-run:
	@if [ -z "$(DAY)" ]; then echo "Usage: make python-run DAY=5"; exit 1; fi
	@printf -v day "%02d" $(DAY)
	@echo "Running Python day $(DAY)..."
	cd python/day$${day} && python solution.py

python-setup:
	@echo "Setting up Python virtual environment..."
	cd python && python3 -m venv venv
	@echo "Installing dependencies..."
	cd python && ./venv/bin/pip install -r requirements.txt
	@echo "Done! Activate with: source python/venv/bin/activate"

# Multi-language targets
test-all: go-test python-test
	@echo "All tests completed!"

clean:
	@echo "Cleaning build artifacts..."
	rm -rf go/**/*.test
	rm -rf python/__pycache__ python/**/__pycache__
	rm -rf python/.pytest_cache
	rm -rf python/**/.pytest_cache
	@echo "Clean complete!"
