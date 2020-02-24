PROJECTNAME=$(shell basename "$(PWD)")
GOLANGCI := $(GOPATH)/bin/golangci-lint

CENT_EMITTER_ADDR?=0x1
CENT_CHAIN_ID?=0x1
CENT_TO?=0x1234567890
CENT_TOKEN_ID?=0x5
CENT_METADATA?=0x0

.PHONY: help run build install
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
	./bin/golangci-lint run ./... --timeout 5m0s

fmt:
	@echo "  >  \033[32mFormatting project...\033[0m "
	gofmt -s -w .

build:
	@echo "  >  \033[32mBuilding binary...\033[0m "
	cd cmd/chainbridge && env GOARCH=amd64 go build -o ../../build/chainbridge

run: build
	@echo "  >  \033[32mRunning bridge...\033[0m "
	./build/chainbridge

install:
	@echo "  >  \033[32mInstalling bridge...\033[0m "
	cd cmd/chainbridge && go install

test:
	@echo "  >  \033[32mRunning tests...\033[0m "
	./scripts/test.sh
	
truffle_test:
	@echo " > \033[32mTesting evm contracts... \033[0m "
	@cd on-chain/evm-contracts && node_modules/.bin/truffle test

truffle_compile:
	@echo " > \033[32mCompiling evm contracts... \033[0m "
	./scripts/evm/compile.sh

start_eth:
	@echo " > \033[32mStarting ganache... \033[0m "
	./scripts/local_test/start_ganache.sh

deploy_eth:
	@echo " > \033[32mDeploying evm contracts... \033[0m "
	./scripts/local_test/ethereum_deploy.sh

docker_start:
	./scripts/docker/start-docker.sh
	
start_cent:
	@echo " > \033[32mStarting centrifuge-chain... \033[0m "
	./scripts/centrifuge/run_chain.sh

cent_auto_run:
	@echo " > \033[32mExecuting centrifuge setup... \033[0m "
	./scripts/centrifuge/run_setup.sh

cent_set_emitter:
	@echo " > \033[32mSetting centrifuge emitter address... \033[0m "
	./scripts/centrifuge/run_interaction.sh set-emitter $(CENT_EMITTER_ADDR)

cent_get_emitter:
	@echo " > \033[32mGetting centrifuge emitter address... \033[0m "
	./scripts/centrifuge/run_interaction.sh get-emitter

cent_whitelist_chain:
	@echo " > \033[32mExecuting centrifuge chain whitelist... \033[0m "
	./scripts/centrifuge/run_interaction.sh whitelist-chain $(CENT_CHAIN_ID)

cent_asset_tx:
	@echo " > \033[32mExecuting centrifuge interaction... \033[0m "
	./scripts/centrifuge/run_interaction.sh asset-tx $(CENT_CHAIN_ID) $(CENT_TO) $(CENT_TOKEN_ID) $(CENT_METADATA)

bindings:
	@echo " > \033[32mCreating go bindings for ethereum contracts... \033[0m "
	./scripts/evm/create_bindings.sh