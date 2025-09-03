# Variables
GO ?= go
PKG := ./...

# Default target
all: test

# Run all tests
test:
	@echo "Running tests..."
	$(GO) test -v $(PKG)

# Run tests with race detector
race:
	@echo "Running tests with race detector..."
	$(GO) test -race -v $(PKG)

# Run tests with coverage
coverage:
	@echo "Running tests with coverage..."
	$(GO) test -coverprofile=coverage.out $(PKG)
	$(GO) tool cover -html=coverage.out

# Run benchmarks
bench:
	@echo "Running benchmarks..."
	$(GO) test -bench=. -benchmem $(PKG)

# Format code
fmt:
	@echo "Formatting code..."
	$(GO) fmt $(PKG)

# Lint code (requires golangci-lint)
lint:
	@echo "Running linter..."
	golangci-lint run

# Build binary (optional)
build:
	@echo "Building binary..."
	$(GO) build -o $(BINARY) main.go

# Clean generated files
clean:
	@echo "Cleaning..."
	rm -f $(BINARY) coverage.out

# Help
help:
	@echo "Makefile targets:"
	@echo "  all        - Run tests (default)"
	@echo "  test       - Run tests"
	@echo "  race       - Run tests with race detector"
	@echo "  coverage   - Run tests with coverage report"
	@echo "  bench      - Run benchmarks"
	@echo "  fmt        - Format code"
	@echo "  lint       - Run linter (golangci-lint)"
	@echo "  build      - Build binary"
	@echo "  clean      - Clean generated files"