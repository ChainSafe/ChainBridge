#!/usr/bin/env bash

# Compile the contracts
sh ./scripts/evm/compile.sh

# Deploy contracts
echo "\nDeploying contracts..."
node ./contracts/evm-contracts/scripts/deploy_local.js --validators 3
