#!/usr/bin/env bash

# Exit on failure
set -e

# Compile the contracts
./scripts/evm/compile.sh

# Deploy contracts
echo "Deploying contracts..."
node ./contracts/evm-contracts/scripts/deploy_local.js --validators 3
