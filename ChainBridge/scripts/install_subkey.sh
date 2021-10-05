#!/usr/bin/env bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only


# Exit on failure
set -e

# Setup cargo & rustc with nightly toolchain
curl https://sh.rustup.rs -sSf | sh -s -- --default-toolchain nightly -y
source $HOME/.cargo/env
rustup target add wasm32-unknown-unknown --toolchain nightly

# Clone and install subkey
git clone --depth 1 https://github.com/paritytech/substrate/
cd substrate
cargo install --force --path ./bin/utils/subkey subkey