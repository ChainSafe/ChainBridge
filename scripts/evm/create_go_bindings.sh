#!/usr/bin/env bash

set -e

contract_path="on-chain/evm-contracts/contracts"
contract_binding_path="on-chain/evm-contracts/"
contract_json_path=./on-chain/evm-contracts/contracts

BIN_DIR=build/bindings/bin
ABI_DIR=build/bindings/abi
GO_DIR=build/bindings/go

mkdir ./build/bindings
mkdir ./build/bindings/go

cp -r $contract_binding_path/bindings/abi $ABI_DIR
cp -r $contract_binding_path/bindings/bin $BIN_DIR

for file in "$BIN_DIR"/*.bin
do
    base=`basename $file`
    value="${base%.*}"
    echo Compiling file $value from path $file

    # Create the go package directory
    mkdir ./contracts/$value
    
    # Build the go package
    abigen --abi $ABI_DIR/${value}.abi --pkg $value --type $value --bin $BIN_DIR/${value}.bin --out $GO_DIR/$value/$value.go
done

rm -rf ./contracts
mkdir ./contracts
cp -r $GO_DIR ./contracts