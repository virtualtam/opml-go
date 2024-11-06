BUILD_DIR ?= build
SRC_FILES := $(shell find . -name "*.go")

all: build lint race cover
.PHONY: all

build: $(BUILD_DIR)/opml2json $(BUILD_DIR)/roundtrip

$(BUILD_DIR)/%: $(SRC_FILES)
	go build -trimpath -o $@ ./cmd/$*

lint:
	golangci-lint run ./...
.PHONY: lint

cover:
	go test -coverprofile=coverage.out ./...
.PHONY: cover

coverhtml: cover
	go tool cover -html=coverage.out
.PHONY: coverhtml

race:
	go test -race ./...
.PHONY: race

test:
	go test ./...
.PHONY: test

# Install development tools
dev-install-tools:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.61.0
	go install github.com/hashicorp/copywrite@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
.PHONY: dev-install-tools

# Licence headers
copywrite:
	copywrite headers
.PHONY: copywrite

# Vulnerability check
vulncheck:
	govulncheck -C . ./...
.PHONY: vulncheck
