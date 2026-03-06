VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
LDFLAGS := -X main.version=$(VERSION)
GOFILES := $(shell find . -name '*.go' -not -path './vendor/*')

.PHONY: build ci fmt-check lint mod-tidy-check test tools vet vuln

build:
	mkdir -p bin
	go build -trimpath -ldflags "$(LDFLAGS)" -o bin/geek-life ./app

fmt-check:
	test -z "$$(gofmt -l $(GOFILES))"

mod-tidy-check:
	go mod tidy -diff

test:
	go test ./...

vet:
	go vet ./...

tools:
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest

lint:
	$$(go env GOPATH)/bin/staticcheck ./...

vuln:
	$$(go env GOPATH)/bin/govulncheck ./...

ci: fmt-check mod-tidy-check vet test lint vuln build
