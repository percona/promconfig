default: help

help:                 ## Display this help message.
	@echo "Please use \`make <target>\` where <target> is one of:"
	@grep '^[a-zA-Z]' $(MAKEFILE_LIST) | \
		awk -F ':.*?## ' 'NF==2 {printf "  %-26s%s\n", $$1, $$2}'


format:                         ## Format source code.
	gofmt -w -s .
	goimports -local github.com/Percona-Lab/promconfig -l -w .

ci:
	go clean -testcache
	go build ./...
	go mod tidy
	git diff --exit-code
