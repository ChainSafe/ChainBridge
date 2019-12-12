PROJECTNAME=$(shell basename "$(PWD)")
GOLANGCI := $(GOPATH)/bin/golangci-lint

.PHONY: help
all: help

help: Makefile
	@echo
	@echo "Choose a make command to run in "$(PROJECTNAME)":"
	@echo
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'
	@echo

get:
	@echo "  >  \033[32mDownloading & Installing all the modules...\033[0m "
	go mod tidy && go mod download

get_lint:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s latest

.PHONY: lint
lint:
	if [ ! -f ./bin/golangci-lint ]; then \
		$(MAKE) get_lint; \
	fi;
	./bin/golangci-lint run

fmt:
	@echo "  >  \033[32mFormatting project...\033[0m "
	gofmt -s -w .

build:
	@echo "  >  \033[32mBuilding binary...\033[0m "
	cd cmd/chainbridge && env GOARCH=amd64 go build -o ./bridge

run:
	@echo "  >  \033[32mRunning main.go...\033[0m "
	go run -v cmd/chainbridge/main.go

test:
	@echo "  >  \033[32mRunning tests...\033[0m "
	./scripts/test.sh