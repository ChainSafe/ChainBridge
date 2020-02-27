#!/usr/bin/env bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only


CMD="go run ."

EMITTER_ADDR=1234
CHAIN_ID=1
TO=0x1234567890
TOKEN_ID=5
METADATA=0

# Exit on failure
set -e

pushd ./scripts/centrifuge/cent-interactions

$CMD set-emitter $EMITTER_ADDR

$CMD whitelist-chain $CHAIN_ID

$CMD asset-tx $CHAIN_ID $TO $TOKEN_ID $METADATA

popd