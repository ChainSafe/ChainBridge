#!/usr/bin/env bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only

DIR="./build/centrifuge-chain"
CENT_CHAIN_URL="https://github.com/ansermino/centrifuge-chain.git"
CENT_CHAIN_BRANCH="david/bridge-module"

# Exit on failure
set -e

rm -rf $DIR
git clone $CENT_CHAIN_URL -b $CENT_CHAIN_BRANCH $DIR
cd $DIR
cargo build --release
cd ..