#!/usr/bin/env bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only

CONTRACTS_REPO="https://github.com/ChainSafe/chainbridge-solidity"
CONTRACTS_BRANCH="wyatt/migrate-pr-200"
CONTRACTS_COMMIT="a4add24d01d3c595a6bfc3c80c300126eb9a7a3d"
CONTRACTS_DIR="./solidity"
DEST_DIR="./bindings"

set -e

git clone -b $CONTRACTS_BRANCH $CONTRACTS_REPO $CONTRACTS_DIR
pushd $CONTRACTS_DIR
git checkout $CONTRACTS_COMMIT

make install-deps
make bindings

popd

mkdir $DEST_DIR
cp -r $CONTRACTS_DIR/build/bindings/go/* $DEST_DIR