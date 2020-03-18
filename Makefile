default: help

help:           ## Display this help message.
	@echo "Please use \`make <target>\` where <target> is one of:"
	@grep '^[a-zA-Z]' $(MAKEFILE_LIST) | \
		awk -F ':.*?## ' 'NF==2 {printf "  %-26s%s\n", $$1, $$2}'

format:         ## Format source code.
	gofmt -w -s .
	goimports -local github.com/percona/promconfig -l -w .

install:        ## Check license and install.
	go run .github/check-license.go
	go install -v ./...

ci: install     ## CI checks.
	git diff --exit-code
