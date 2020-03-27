#!/usr/bin/env bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only

CONTRACTS_REPO="https://github.com/ChainSafe/chainbridge-solidity"
CONTRACTS_BRANCH="test/threshold-increase"
CONTRACTS_COMMIT="b3feb0f2bc9ded0194e84315bb9b910ca03a9934"
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