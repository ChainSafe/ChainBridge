#!/usr/bin/env bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only

S3_URL="https://centchain.nyc3.digitaloceanspaces.com"
SUB_COMMIT="4bf5a2b7b3cba16a9c1c0ab93115fd4886502733"
SUB_BUILD_ID="155080360"
SUB_CMD="chainbridge-substrate-chain"

set -eux

rm -rf ./$SUB_CMD

wget $S3_URL/$SUB_CMD/$SUB_COMMIT-$CENT_BUILD_ID/$SUB_CMD

chmod a+x ./$SUB_CMD

rm -rf $HOME/.local/share/$SUB_CMD/

./$SUB_CMD --dev --alice &