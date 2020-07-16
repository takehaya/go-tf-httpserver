NAME := go-tf-httpserver

# tag name version
# VERSION := $(shell git describe --tags --abbrev=0)

#brunch name version
VERSION := $(shell git rev-parse --abbrev-ref HEAD)

PKG_NAME=$(shell basename `pwd`)

LDFLAGS := -ldflags="-s -w  -X \"https://github.com/takehaya/go-tf-httpserver/pkg/version.Version=$(VERSION)\" -extldflags \"-static\""
SRCS    := $(shell find . -type f -name '*.go')

.DEFAULT_GOAL := build
build: $(SRCS)
	go build $(LDFLAGS) -o ./bin/$(NAME) ./cmd/go-tf-httpserver

.PHONY: run
run:
	go run $(LDFLAGS) ./cmd/go-tf-httpserver

.PHONY: test
test:
	go test -v ./integration

## lint
.PHONY: lint
lint:
	@for pkg in $$(go list ./...): do \
		golint --set_exit_status $$pkg || exit $$?; \
	done

.PHONY: clean
clean:
	rm -rf bin

.PHONY: install
install:
	go install $(LDFLAGS) ./cmd/go-tf-httpserver
