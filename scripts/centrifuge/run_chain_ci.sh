#!/usr/bin/env bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only

S3_URL="https://centchain.nyc3.digitaloceanspaces.com"
CENT_CHAIN_COMMIT="4bf5a2b7b3cba16a9c1c0ab93115fd4886502733"
CENT_CHAIN_BUILD_ID="155080360"

set -eux

rm -rf ./centrifuge-chain

wget $S3_URL/$CENT_CHAIN_COMMIT-$CENT_CHAIN_BUILD_ID/centrifuge-chain

chmod a+x ./centrifuge-chain

rm -rf $HOME/.local/share/centrifuge-chain/

./centrifuge-chain --dev --alice &