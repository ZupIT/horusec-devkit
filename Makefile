GO ?= go
GOFMT ?= gofmt
GO_FILES ?= $$(find . -name '*.go' | grep -v vendor)
GO_CI_LINT ?= ./bin/golangci-lint
GO_IMPORTS ?= goimports
GO_IMPORTS_LOCAL ?= github.com/ZupIT/horusec-devkit

fmt:
	$(GOFMT) -w $(GO_FILES)

lint:
    ifeq ($(wildcard $(GO_CI_LINT)), $(GO_CI_LINT))
		$(GO_CI_LINT) run -v --timeout=5m -c .golangci.yml ./...
    else
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s latest
		$(GO_CI_LINT) run -v --timeout=5m -c .golangci.yml ./...
    endif

coverage:
	chmod +x scripts/coverage.sh
	scripts/coverage.sh 90 "."

test:
	$(GO) clean -testcache && $(GO) test -v ./... -timeout=20m -parallel=1 -failfast -short

imports:
    ifeq (, $(shell which $(GO_IMPORTS)))
		$(GO) get -u golang.org/x/tools/cmd/goimports
		$(GO_IMPORTS) -local $(GO_IMPORTS_LOCAL) -w $(GO_FILES)
    else
		$(GO_IMPORTS) -local $(GO_IMPORTS_LOCAL) -w $(GO_FILES)
    endif
	
pipeline: fmt imports lint test coverage
