# Makefile for httpgo

# Variables
BINARY_NAME=httpgo
PKG_PATH=httpgo/pkg/http/api

# Default version, can be overridden
VERSION ?= dev
BUILD_TIME := $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
# Attempt to get commit hash, default to 'unknown' if not a git repo or git is not found
COMMIT_HASH := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

# Go parameters
GO=go
GO_BUILD=$(GO) build
GO_CLEAN=$(GO) clean

# Linker flags
LDFLAGS = -ldflags="-X $(PKG_PATH).Version=$(VERSION) -X '$(PKG_PATH).BuildTime=$(BUILD_TIME)' -X $(PKG_PATH).CommitHash=$(COMMIT_HASH)"

# Default target
all: build

# Build the binary
build:
	@echo "Building $(BINARY_NAME) with version $(VERSION)..."
	$(GO_BUILD) $(LDFLAGS) -o $(BINARY_NAME) main.go
	@echo "$(BINARY_NAME) built successfully."

# Run the binary
run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BINARY_NAME)

# Clean the binary
clean:
	@echo "Cleaning..."
	$(GO_CLEAN)
	rm -f $(BINARY_NAME)
	@echo "Clean complete."

# Show version information (by building and running a temporary binary that just prints version)
show_version:
	@echo "Version: $(VERSION)"
	@echo "Build Time: $(BUILD_TIME)"
	@echo "Commit Hash: $(COMMIT_HASH)"

.PHONY: all build run clean show_version
