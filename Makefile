# Variables
GO ?= go
PKG := ./...
BENCH ?= .
RUN ?= .
MEM ?= -benchmem
TIME ?= 5s
COUNT ?= 1
CPU ?= 8
TAGS ?=
EXTRA ?=
ITER ?=
TIMEOUT ?= 10m
VERBOSE ?= false
CPU_PROFILE ?= cpu.out
MEM_PROFILE ?= mem.out

.PHONY: all test race coverage bench bench-profile bench-all-profiles flamegraph fmt lint clean help

# Default target
all: test

# Run all tests
test:
	@echo "Running tests for package $(PKG)"
	$(GO) test -run $(RUN) $(PKG) -v

# Run tests with race detector
race:
	@echo "Running tests with race detector..."
	$(GO) test -race -v $(PKG)

# Run tests with coverage
coverage:
	@echo "Running tests with coverage..."
	$(GO) test -coverprofile=coverage.out $(PKG)
	$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "✅ Coverage report: coverage.html"

# Run benchmarks
bench:
	@echo "Running benchmarks for package: $(PKG) with BENCH='$(BENCH)'"
	$(GO) test -run $(RUN) -bench '$(BENCH)' $(MEM) -benchtime=$(TIME) $(ITER) \
	-count=$(COUNT) -cpu $(CPU) -tags '$(TAGS)' -timeout $(TIMEOUT) $(EXTRA) $(PKG)

# Profiled benchmarks (single package only)
bench-profile:
	@if [ "$(PKG)" = "./..." ]; then \
		echo "❌ Error: Cannot use bench-profile with multiple packages (./...). Set PKG to a single package (e.g., ./treemap)"; \
		exit 1; \
	fi
	@echo "Running profiled benchmarks:"
	@echo "  Package: $(PKG)"
	@echo "  Benchmark: $(BENCH)"
	@echo "  Time: $(TIME), Count: $(COUNT), CPU: $(CPU)"
	$(GO) test -run $(RUN) -bench '$(BENCH)' $(MEM) -benchtime=$(TIME) $(ITER) \
	-count=$(COUNT) -cpuprofile $(CPU_PROFILE) -memprofile $(MEM_PROFILE) \
	-cpu $(CPU) -tags '$(TAGS)' -timeout $(TIMEOUT) $(EXTRA) $(PKG)
	@echo "✅ CPU profile saved to $(CPU_PROFILE)"
	@echo "✅ Memory profile saved to $(MEM_PROFILE)"

# Profile all packages (each gets separate files)
bench-all-profiles:
	@echo "Running profiled benchmarks for ALL packages..."
	@for pkg in $$(go list ./...); do \
		f=$$(echo $$pkg | tr / _); \
		echo "Benchmarking $$pkg..."; \
		$(GO) test -bench $(BENCH) $(MEM) -benchtime=$(TIME) $(ITER) \
		-count=$(COUNT) -cpuprofile "$${f}_cpu.out" \
		-memprofile "$${f}_mem.out" \
		-cpu $(CPU) -tags '$(TAGS)' -timeout $(TIMEOUT) $(EXTRA) $$pkg; \
	done

# Generate flamegraph using go tool pprof (interactive)
flamegraph:
	@echo "Generating flamegraph from $(CPU_PROFILE)..."
	go tool pprof -http=:8080 $(CPU_PROFILE)

# Format code
fmt:
	@echo "Formatting code..."
	$(GO) fmt $(PKG)

# Lint code (requires golangci-lint)
lint:
	@if ! command -v golangci-lint >/dev/null 2>&1; then \
		echo "❌ golangci-lint not installed. Install: https://golangci-lint.run/usage/install/"; \
		exit 1; \
	fi
	@echo "Running linter..."
	golangci-lint run

# Clean generated files (profiles, coverage, binaries)
clean:
	@echo "Cleaning generated files..."
	@rm -f $(CPU_PROFILE) $(MEM_PROFILE) coverage.out coverage.html *.test
	@rm -f *_cpu.out *_mem.out
	@echo "✅ Clean complete."

# Help
help:
	@echo "Makefile targets:"
	@echo "  all            - Run tests (default)"
	@echo "  test           - Run tests"
	@echo "  race           - Run tests with race detector"
	@echo "  coverage       - Run tests with coverage report"
	@echo "  bench          - Run benchmarks (customize with PKG, BENCH, TIME, COUNT, etc.)"
	@echo "  bench-profile  - Run benchmarks with CPU/mem profiles (cpu.out, mem.out)"
	@echo "  bench-all-profiles - Profile all packages"
	@echo "  flamegraph     - Open interactive flamegraph on :8080"
	@echo "  fmt            - Format code"
	@echo "  lint           - Run linter (golangci-lint)"
	@echo "  clean          - Clean generated files"