#!/usr/bin/env bash

contract_path="on-chain/evm-contracts/contracts"
contract_binding_path="on-chain/evm-contracts/"
contracts="Receiver BridgeAsset Emitter Safe SimpleEmitter"

BIN_DIR=build/bindings/bin
ABI_DIR=build/bindings/abi
GO_DIR=build/bindings/go

rm -rf ./build/bindings
rm -rf ./build/bindings/go
rm -rf ./contracts

mkdir ./build/bindings
mkdir ./build/bindings/go
mkdir ./contracts

cp -r $contract_binding_path/bindings/abi $ABI_DIR
cp -r $contract_binding_path/bindings/bin $BIN_DIR

contract_json_path=./on-chain/evm-contracts/contracts

for file in "$BIN_DIR"/*.bin
do
    base=`basename $file`
    value="${base%.*}"
    echo Compiling file $value from path $file

    # Create the go package directory
    mkdir ./contracts/$value
    
    # Build the go package
    abigen --abi $ABI_DIR/${value}.abi --pkg $value --type $value --bin $BIN_DIR/${value}.bin --out contracts/$value/$value.go
done