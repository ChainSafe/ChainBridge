#!/usr/bin/env bash

contract_path="./contracts/evm-contracts/"
build_dir="./build/evm-contracts/"

echo "Compiling contracts..."

pushd $contract_path

npm install --silent
node_modules/.bin/truffle compile

popd

echo "Copying contracts to root..."
mkdir -p $build_dir && cp -r $contract_path/build/contracts/ $build_dir
