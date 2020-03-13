#!/usr/bin/env bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only


# Exit on failure
set -e

PORT=${PORT:-8545}

# Compile the contracts
./scripts/evm/compile.sh

# Deploy contracts
echo "Deploying contracts..."
node ./on-chain/evm-contracts/scripts/cli/index.js --validators 3 --port $PORT
