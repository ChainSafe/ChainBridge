#!/usr/bin/env bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only

set -eux

# TODO: Clone repo once #322 is closed
# TODO: Pass commit from Makefile
# git clone https://github.com/ChainSafe/chainbridge-solidity && cd chainbridge-solidity
# git checkout $SOL_COMMIT

cd solidity/
make install-cli