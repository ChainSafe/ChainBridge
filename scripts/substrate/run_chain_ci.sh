#!/usr/bin/env bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only

S3_URL="https://centchain.nyc3.digitaloceanspaces.com"
SUB_COMMIT="911fa5a56bda6de383960905e1cc14ae41dd0267"
SUB_BUILD_ID="156367520"
SUB_CMD="chainbridge-substrate-chain"

set -eux

rm -rf ./$SUB_CMD

wget $S3_URL/$SUB_CMD/$SUB_COMMIT-$SUB_BUILD_ID/$SUB_CMD

chmod a+x ./$SUB_CMD

rm -rf $HOME/.local/share/$SUB_CMD/

./$SUB_CMD --dev --alice &