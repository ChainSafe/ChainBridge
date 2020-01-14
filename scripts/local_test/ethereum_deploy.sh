#!/usr/bin/env bash

# Exit on failure
set -e

PORT=${PORT:-8545}

# Compile the contracts
./scripts/evm/compile.sh

# Deploy contracts
echo "Deploying contracts..."
node ./contracts/evm-contracts/scripts/deploy_local.js --validators 3 --port $PORT
