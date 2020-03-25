#!/usr/bin/env bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only

CONTRACTS_REPO="https://github.com/ChainSafe/chainbridge-solidity"
CONTRACTS_DIR="./solidity"
DEST_DIR="./bindings"

set -e

git clone $CONTRACTS_REPO $CONTRACTS_DIR

pushd $CONTRACTS_DIR

make install-deps
make bindings

popd

mkdir $DEST_DIR
cp -r $CONTRACTS_DIR/build/bindings/go/* $DEST_DIR