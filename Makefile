# Project information
BINARY_NAME=go-bootstrap
BUILD_DIR=./build
MAIN_PACKAGE=./cmd/go-bootstrap
VERSION=0.0.2-alpha
LDFLAGS=-ldflags "-X github.com/paoloanzn/go-bootstrap/config.VERSION=$(VERSION)"

# Default installation paths
PREFIX=/usr/local
DESTDIR=

# Go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

.PHONY: all build clean test run mod-tidy help install dev

# Default target
all: build

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	@$(GOCLEAN)
	@rm -rf $(BUILD_DIR)

# Run tests
test:
	@echo "Running tests..."
	@$(GOTEST) -v ./...

# Build and run
run: build
	@echo "Running $(BINARY_NAME)..."
	@$(BUILD_DIR)/$(BINARY_NAME) $(ARGS)

# Clean up dependencies
mod-tidy:
	@echo "Tidying Go modules..."
	@$(GOCMD) mod tidy

# Install binary
install: build
	@echo "Installing $(BINARY_NAME)..."
	@mkdir -p $(DESTDIR)$(PREFIX)/bin
	@install -m 755 $(BUILD_DIR)/$(BINARY_NAME) $(DESTDIR)$(PREFIX)/bin/$(BINARY_NAME)

# Cross-compile for multiple platforms
cross-build:
	@echo "Cross-compiling for multiple platforms..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_PACKAGE)
	@GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(MAIN_PACKAGE)
	@GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN_PACKAGE)

# Format code
fmt:
	@echo "Formatting code..."
	@$(GOCMD) fmt ./...

# Vet code for potential issues
vet:
	@echo "Vetting code..."
	@$(GOCMD) vet ./...

# Development workflow: format, vet, build
dev: fmt vet build

# Help information
help:
	@echo "Available targets:"
	@echo "  build       - Build the binary"
	@echo "  clean       - Remove build artifacts"
	@echo "  test        - Run tests"
	@echo "  run         - Build and execute (use ARGS='start -dev' for arguments)"
	@echo "  mod-tidy    - Clean up dependencies"
	@echo "  install     - Install binary to system location (default: /usr/local/bin)"
	@echo "  cross-build - Build for multiple platforms (Linux, macOS, Windows)"
	@echo "  fmt         - Format code"
	@echo "  vet         - Run Go vet to find potential issues"
	@echo "  dev         - Run format, vet, and build"