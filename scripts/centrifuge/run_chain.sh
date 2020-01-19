#!/usr/bin/env bash

# Exit on failure
set -e

# Temporary repository until stable
# Tailored image to not validate anchor or proofs
docker run -p 9933:9933 -p 9944:9944 -p 30333:30333 mikiquantum/centrifuge-chain:latest centrifuge-chain --dev --ws-external --rpc-external
