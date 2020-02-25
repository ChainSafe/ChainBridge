#!/usr/bin/env bash

set -e

contract_binding_path="on-chain/evm-contracts/"
base_path="./build/tmp"
BIN_DIR="$base_path/bin"
ABI_DIR="$base_path/abi"
GO_DIR="$base_path/go"

# Remove old bin and abi
echo "Removing old builds..."
rm -rf $base_path
mkdir $base_path
mkdir $ABI_DIR
mkdir $BIN_DIR
mkdir $GO_DIR

echo "Copying new builds to root..."
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

# Remove old bindings
rm -rf ./contracts
mkdir ./contracts

# Copy in new bindings
cp -r $GO_DIR ./contracts

# cleanup tmp
rm -rf ./build/tmp