default: help

help:                    ## Display this help message.
	@echo "Please use \`make <target>\` where <target> is one of:"
	@grep '^[a-zA-Z]' $(MAKEFILE_LIST) | \
		awk -F ':.*?## ' 'NF==2 {printf "  %-26s%s\n", $$1, $$2}'

init:                    ## Install development tools
	cd tools && go generate -x -tags=tools

FILES = $(shell find . -type f -name '*.go')

format:                  ## Format source code.
	bin/gofumpt -l -w $(FILES)
	bin/goimports -local github.com/percona/pmm -l -w $(FILES)

check:                   ## Run checks/linters for the whole project
	go run .github/check-license.go
	bin/go-consistent -pedantic ./...
	bin/golangci-lint run

test:                    ## Run tests
	go test -race -timeout=10m ./...

test-cover:              ## Run tests and collect per-package coverage information
	go test -race -timeout=10m -count=1 -coverprofile=cover.out -covermode=atomic ./...

test-crosscover:         ## Run tests and collect cross-package coverage information
	go test -race -timeout=10m -count=1 -coverprofile=crosscover.out -covermode=atomic -p=1 -coverpkg=./... ./...

ci: install              ## CI checks.
	git diff --exit-code
