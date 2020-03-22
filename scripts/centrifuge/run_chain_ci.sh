#!/usr/bin/env bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only

S3_URL="https://centchain.nyc3.digitaloceanspaces.com"
CENT_CHAIN_COMMIT="6d1a3f1641513ae211ad46f80e8cd739eb6a0619"
CENT_CHAIN_BUILD_ID="154078537"
CENT_CHAIN_BRANCH="david/add-bridge"

set -eux

rm -rf ./centrifuge-chain

wget $S3_URL/$CENT_CHAIN_COMMIT-$CENT_CHAIN_BUILD_ID/centrifuge-chain

chmod a+x ./centrifuge-chain

rm -rf $HOME/.local/share/centrifuge-chain/

./centrifuge-chain --dev --alice &