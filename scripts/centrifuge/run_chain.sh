#!/usr/bin/env bash

# Exit on failure
set -e

# If subkey command not found
if ! command -v subkey 2>/dev/null
then
  curl https://getsubstrate.io -sSf | bash -s -- --fast
  cargo install --force --git https://github.com/paritytech/substrate subkey
fi

# Temporary repository until stable
# Tailored image to not validate anchor or proofs
docker run -p 9933:9933 -p 9944:9944 -p 30333:30333 mikiquantum/centrifuge-chain:latest centrifuge-chain --dev --ws-external --rpc-external
