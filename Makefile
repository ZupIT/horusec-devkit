GO ?= go
GOFMT ?= gofmt
GO_FILES ?= $$(find . -name '*.go' | grep -v vendor)
GOLANG_CI_LINT ?= golangci-lint
GO_IMPORTS ?= goimports
GO_FUMPT ?= gofumpt
GO_GCI ?= gci
GO_IMPORTS_LOCAL ?= github.com/ZupIT/horusec-devkit
HORUSEC ?= horusec
COMPOSE_FILE_NAME ?= docker-compose.yaml
DOCKER_COMPOSE ?= docker-compose
ADDLICENSE ?= addlicense

lint:
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.43.0
	$(GOLANG_CI_LINT) run -v --timeout=5m -c .golangci.yml ./...

coverage:
	chmod +x scripts/coverage.sh
	scripts/coverage.sh 99.4 "."

test:
	$(GO) clean -testcache && $(GO) test -v ./... -timeout=2m -parallel=1 -failfast -short

install-format-dependencies:
	$(GO) install golang.org/x/tools/cmd/goimports@v0.1.7
	$(GO) install mvdan.cc/gofumpt@v0.2.0
	$(GO) install github.com/daixiang0/gci@v0.2.9

format: install-format-dependencies
	$(GOFMT) -s -l -w $(GO_FILES)
	$(GO_IMPORTS) -w -local $(GO_IMPORTS_LOCAL) $(GO_FILES)
	$(GO_FUMPT) -l -w $(GO_FILES)
	$(GO_GCI) -w -local $(GO_IMPORTS_LOCAL) $(GO_FILES)

security:
    ifeq (, $(shell which $(HORUSEC)))
		curl -fsSL https://raw.githubusercontent.com/ZupIT/horusec/master/deployments/scripts/install.sh | bash -s latest
		$(HORUSEC) start -p="./" -e="true"
    else
		$(HORUSEC) start -p="./" -e="true"
    endif

migrate-up:
	chmod +x ./scripts/migration-run.sh
	./scripts/migration-run.sh up

migrate-drop:
	chmod +x ./scripts/migration-run.sh
	./scripts/migration-run.sh drop -f

update-auth-grpc:
	protoc --go_out=.  --go-grpc_out=.  ./pkg/services/grpc/auth/proto/auth.proto

license:
	$(GO) install github.com/google/addlicense@v1.0.0
	@$(ADDLICENSE) -check -f ./copyright.txt $(shell find -regex '.*\.\(go\|js\|ts\|yml\|yaml\|sh\|dockerfile\)')

license-fix:
	$(GO) install github.com/google/addlicense@v1.0.0
	@$(ADDLICENSE) -f ./copyright.txt $(shell find -regex '.*\.\(go\|js\|ts\|yml\|yaml\|sh\|dockerfile\)')

pipeline: format lint test coverage security

migrate: migrate-drop migrate-up

