GO ?= go
GOFMT ?= gofmt
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
GO_CI_LINT ?= ./bin/golangci-lint

fmt:
	$(GOFMT) -w $(GOFMT_FILES)

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

pipeline: fmt lint test coverage