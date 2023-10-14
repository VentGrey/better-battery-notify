# Build Variables
BIN_NAME = "battery-notify"
GOPATH ?= $(shell go env GOPATH)
GOLDFLAGS ?= "-s -w"

# Default target
all: build

build:
	@echo "==> Building $(BIN_NAME)..."
	CGO_ENABLED=0 go build -o $(BIN_NAME) -ldflags $(GOLDFLAGS) .

clean:
	@echo "==> Cleaning..."
	@rm -f $(BIN_NAME)
	go clean

clean-deps:
	@echo "==> Cleaning go dependencies cache..."
	go clean -modcache

test:
	@echo "==> Running tests..."
	go test -v

ci: clean build test
	@echo "==> CI pipeline completed successfully."

help:
	@echo "Available targets:"
	@echo "  all (default)    - Build the project"
	@echo "  build            - Build the project"
	@echo "  clean            - Remove built binaries"
	@echo "  clean-deps       - Remove go dependencies cache"
	@echo "  test             - Run tests"
	@echo "  ci               - Run CI"
	@echo "  help             - Display this help message"
