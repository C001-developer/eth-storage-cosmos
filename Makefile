# Linting, Teseting
golangci_lint_cmd=github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.2

install-linter:
	@echo "--> Installing linter"
	@go install $(golangci_lint_cmd)

lint:
	@echo "--> Running linter"
	@ $$(go env GOPATH)/bin/golangci-lint run --timeout=10m
.PHONY:	lint install-linter

test:
	go test ./... --timeout=10m
.PHONY: test