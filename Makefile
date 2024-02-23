# Avoid printing directory changes
MAKEFLAGS += --no-print-directory

# Set shell options for safety
.SHELLFLAGS := -e -u -o pipefail -c

# Default target
.PHONY: all
all: help

# Use bash for advanced scripting capabilities, including inline if-statements
SHELL := /bin/bash

# Detect operating system
UNAME_S := $(shell uname -s | awk '{print tolower($$0)}')
BINARY_NAME := interpol

##@ Helpers

# Dynamically generates and displays the help message
help: ## Displays this help message
	@awk 'BEGIN {FS = ":.*##"; printf "\033[36mUsage:\033[0m\n"} /^[a-zA-Z0-9_%\/-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
	@printf "\n"

##@ Compilation
build: ## Cross compile for Darwin and Linux
	@GOOS=darwin GOARCH=arm64 go build -o ./bin/darwin/$(BINARY_NAME) main.go
	@GOOS=darwin GOARCH=arm64 go build -o ./bin/darwin/$(BINARY_NAME)_arm64 main.go
	@GOOS=linux GOARCH=amd64 go build -o ./bin/linux/$(BINARY_NAME) main.go
	@GOOS=linux GOARCH=amd64 go build -o ./bin/linux/$(BINARY_NAME)_amd64 main.go

clean: ## Remove binary files
	@go clean
	@rm -f ./bin/darwin/$(BINARY_NAME)
	@rm -f ./bin/linux/$(BINARY_NAME)

##@ Rendering
# Example: make render/env-brace.json
render-brace/%: export PGPASSWORD=pass1234
render-brace/%: build ## Render file (e.g., make render/env-brace.json)
	@./bin/$(UNAME_S)/$(BINARY_NAME) $(notdir $@)

# Example: make render/env-brace.json
render-square/%: export PGPASSWORD=pass1234
render-square/%: build ## Render file (e.g., make render/env-square.json)
	@./bin/$(UNAME_S)/$(BINARY_NAME) -s $(notdir $@)