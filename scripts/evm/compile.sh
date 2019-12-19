#!/usr/bin/env bash

contract_path="./contracts/evm-contracts/"
build_dir="./build/evm-contracts/"

pushd $contract_path

echo "\nRunning npm install..."
npm install --silent
echo "Running truffle compile..."
node_modules/.bin/truffle compile

popd

echo "\nCopying contracts to root..."
mkdir -p $build_dir && cp -r $contract_path/build/contracts/ $build_dir
