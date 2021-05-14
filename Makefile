PROJECTNAME=$(shell basename "$(PWD)")
VERSION=-ldflags="-X main.Version=$(shell git describe --tags)"
SOL_DIR=./solidity

CENT_EMITTER_ADDR?=0x1
CENT_CHAIN_ID?=0x1
CENT_TO?=0x1234567890
CENT_TOKEN_ID?=0x5
CENT_METADATA?=0x0

.PHONY: help run build install license
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

get-lint:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s v1.31.0

.PHONY: lint
lint:
	if [ ! -f ./bin/golangci-lint ]; then \
		$(MAKE) get-lint; \
	fi;
	./bin/golangci-lint run ./... --timeout 5m0s

lint-fix:
	if [ ! -f ./bin/golangci-lint ]; then \
		$(MAKE) get-lint; \
	fi;
	./bin/golangci-lint run ./... --timeout 5m0s --fix

build:
	@echo "  >  \033[32mBuilding binary...\033[0m "
	cd cmd/chainbridge && env GOARCH=amd64 go build -o ../../build/chainbridge $(VERSION)

install:
	@echo "  >  \033[32mInstalling bridge...\033[0m "
	cd cmd/chainbridge && go install $(VERSION)

build-mkdocs:
	docker run --rm -it -v ${PWD}:/docs squidfunk/mkdocs-material build

setup-sol-cli:
	@echo " > \033[32mSetting up solidity cli... \033[0m "
	TARGET=cli-only ./scripts/setup_contracts.sh

rebuild-contracts:
	rm -rf bindings/ solidity/
	TARGET=build ./scripts/setup_contracts.sh

## license: Adds license header to missing files.
license:
	@echo "  >  \033[32mAdding license headers...\033[0m "
	GO111MODULE=off go get -u github.com/google/addlicense
	addlicense -c "ChainSafe Systems" -f ./scripts/header.txt -y 2020 .

## license-check: Checks for missing license headers
license-check:
	@echo "  >  \033[Checking for license headers...\033[0m "
	GO111MODULE=off go get -u github.com/google/addlicense
	addlicense -check -c "ChainSafe Systems" -f ./scripts/header.txt -y 2020 .

## Install dependency subkey
install-subkey:
	curl https://getsubstrate.io -sSf | bash -s -- --fast
	cargo install --force --git https://github.com/paritytech/substrate subkey

## Runs go test for all packages except the solidity bindings
test:
	@echo "  >  \033[32mRunning tests...\033[0m "
	go test -p 1 -coverprofile=cover.out -v `go list ./... | grep -v bindings | grep -v e2e`

test-e2e:
	@echo "  >  \033[32mRunning e2e tests...\033[0m "
	go test -p 1 -v -timeout 0 ./e2e

test-eth:
	@echo "  >  \033[32mRunning ethereum tests...\033[0m "
	go test ./chains/ethereum

test-sub:
	@echo "  >  \033[32mRunning substrate tests...\033[0m "
	go test ./chains/substrate

docker-start:
	./scripts/docker/start-docker.sh

docker-e2e:
	docker-compose -f ./docker-compose-e2e.yml up -V

mkdocs:
	docker run --rm -it -p 8000:8000 -v ${PWD}:/docs squidfunk/mkdocs-material

clean:
	rm -rf build/ solidity/
