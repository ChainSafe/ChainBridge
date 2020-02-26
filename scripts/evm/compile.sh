#!/usr/bin/env bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only


contract_path="./on-chain/evm-contracts/"
build_dir="./build/evm-contracts/"

echo "Compiling contracts..."

pushd $contract_path

npm install --silent
node_modules/.bin/truffle compile

popd

echo "Copying contracts to root..."
mkdir -p $build_dir && cp -r $contract_path/build/contracts/ $build_dir
