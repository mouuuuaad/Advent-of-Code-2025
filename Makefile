.PHONY: help go-test go-run go-bench clean

# Default target
help:
	@echo "Advent of Code 2025 - Multi-Language Solutions"
	@echo ""
	@echo "Available targets:"
	@echo "  go-test          - Run all Go tests"
	@echo "  go-run DAY=X     - Run specific Go day (1-4)"
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

clean:
	@echo "Cleaning build artifacts..."
	rm -rf go/**/*.test
	rm -f c/day*/*.o c/day*/solution
	rm -f cpp/day*/*.o cpp/day*/solution
	rm -f odin/day*/*.bin
	@echo "Clean complete!"
