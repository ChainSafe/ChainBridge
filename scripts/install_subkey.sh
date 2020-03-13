#!/usr/bin/env bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only


# Exit on failure
set -e

git clone --depth 1 https://github.com/paritytech/substrate/
cd substrate
cargo +nightly install --force --path ./bin/utils/subkey subkey