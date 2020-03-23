#!/usr/bin/env bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only


set -e

echo Top

# Compile contracts
echo running compile.sh
./scripts/evm/compile.sh

# Switch to directory
pushd ./on-chain/evm-contracts

# Generate bindings
echo running compileAbiBin.js
node ./scripts/compileAbiBin.js

# Go back to root
popd

# Generate go bindings
echo running create_go_bindings.sh
./scripts/evm/create_go_bindings.sh
