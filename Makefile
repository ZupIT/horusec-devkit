GO ?= go
GOFMT ?= gofmt
GO_FILES ?= $$(find . -name '*.go' | grep -v vendor)
GOLANG_CI_LINT ?= ./bin/golangci-lint
GO_IMPORTS ?= goimports
GO_IMPORTS_LOCAL ?= github.com/ZupIT/horusec-devkit
HORUSEC ?= horusec
COMPOSE_FILE_NAME ?= docker-compose.yaml
DOCKER_COMPOSE ?= docker-compose
ADDLICENSE ?= addlicense

fmt:
	$(GOFMT) -w $(GO_FILES)
vet:  ## vet go files
	go vet `go list ./... | grep -v /vendor/`

lint:
    ifeq ($(wildcard $(GOLANG_CI_LINT)), $(GOLANG_CI_LINT))
		$(GOLANG_CI_LINT) run -v --timeout=5m -c .golangci.yml ./...
    else
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s latest
		$(GOLANG_CI_LINT) run -v --timeout=5m -c .golangci.yml ./...
    endif

coverage:
	chmod +x scripts/coverage.sh
	scripts/coverage.sh 99.3 "."

test:
	$(GO) clean -testcache && $(GO) test -v ./... -timeout=2m -parallel=1 -failfast -short

fix-imports:
    ifeq (, $(shell which $(GO_IMPORTS)))
		$(GO) get -u golang.org/x/tools/cmd/goimports
		$(GO_IMPORTS) -local $(GO_IMPORTS_LOCAL) -w $(GO_FILES)
    else
		$(GO_IMPORTS) -local $(GO_IMPORTS_LOCAL) -w $(GO_FILES)
    endif

security:
    ifeq (, $(shell which $(HORUSEC)))
		curl -fsSL https://raw.githubusercontent.com/ZupIT/horusec/master/deployments/scripts/install.sh | bash -s latest
		$(HORUSEC) start -p="./" -e="true"
    else
		$(HORUSEC) start -p="./" -e="true"
    endif

migrate: migrate-drop migrate-up

migrate-up:
	chmod +x ./scripts/migration-run.sh
	./scripts/migration-run.sh up

migrate-drop:
	chmod +x ./scripts/migration-run.sh
	./scripts/migration-run.sh drop -f

update-auth-grpc:
	protoc --go_out=.  --go-grpc_out=.  ./pkg/services/grpc/auth/proto/auth.proto

pipeline: fmt fix-imports lint test coverage security

license:
	$(GO) get -u github.com/google/addlicense
	@$(ADDLICENSE) -check -f ./copyright.txt $(shell find -regex '.*\.\(go\|js\|ts\|yml\|yaml\|sh\|dockerfile\)')

license-fix:
	$(GO) get -u github.com/google/addlicense
	@$(ADDLICENSE) -f ./copyright.txt $(shell find -regex '.*\.\(go\|js\|ts\|yml\|yaml\|sh\|dockerfile\)')
