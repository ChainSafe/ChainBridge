#!/usr/bin/env bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only

CONTRACTS_REPO="https://github.com/ChainSafe/chainbridge-solidity"
CONTRACTS_BRANCH="wyatt/migrate-pr-200"
CONTRACTS_COMMIT="2634b786afe01104516df4c6a5a404dd3007aa7b"
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