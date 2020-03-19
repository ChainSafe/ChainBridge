#!/usr/bin/env bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only

S3_URL="https://centchain.nyc3.digitaloceanspaces.com"
CENT_CHAIN_COMMIT="51f7a9dfbf51590dceb457da0a8a05c814cbabef"
CENT_CHAIN_BUILD_NUM="151844930"
CENT_CHAIN_BRANCH="david/bridge-module"

set -eux

rm -rf ./centrifuge-chain

wget $S3_URL/$CENT_CHAIN_COMMIT-$CENT_CHAIN_BUILD_NUM/centrifuge-chain

chmod a+x ./centrifuge-chain

rm -rf $HOME/.local/share/centrifuge-chain/

./centrifuge-chain --dev --alice &