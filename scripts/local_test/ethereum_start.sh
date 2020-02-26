#!/usr/bin/env bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only


pushd ./on-chain/evm-contracts &&
npm install &&
node ./scripts/deploy_local.js --validators 3
