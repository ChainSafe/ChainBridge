#!/usr/bin/env bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only

CONTRACTS_REPO="https://github.com/ChainSafe/chainbridge-solidity"
CONTRACTS_BRANCH="master"
CONTRACTS_COMMIT="361310aa70ab943ab0a022f6377f71f86135cb0f"
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