default: help

help:			## Display this help message.
	@echo "Please use \`make <target>\` where <target> is one of:"
	@grep '^[a-zA-Z]' $(MAKEFILE_LIST) | \
		awk -F ':.*?## ' 'NF==2 {printf "  %-26s%s\n", $$1, $$2}'

format:			## Format source code.
	gofmt -w -s .
	goimports -local github.com/Percona-Lab/promconfig -l -w .

lint:			## Linter checks
	golangci-lint run

deps:			## Resolve dependencies
	go mod tidy

update:			## Update dependencies
	go get -u
	go mod tidy

ci:				## CI checks
	go clean -testcache
	go build ./...
	make deps
	git diff --exit-code

.PHONY: help format lint deps update ci
