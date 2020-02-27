#!/usr/bin/env bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only


# Exit on failure
set -e

pushd ./scripts/centrifuge/cent-interactions

go run . "$@"

popd