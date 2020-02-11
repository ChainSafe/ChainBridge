#!/usr/bin/env bash

CENT_CHAIN_URL="https://github.com/ansermino/centrifuge-chain.git"
CENT_CHAIN_BRANCH="david/bridge-module"

# Exit on failure
set -e

if [ ! -d "./centrifuge-chain" ]
then
  echo "Centrifuge-chain not found, fetching and building"
  git clone $CENT_CHAIN_URL -b $CENT_CHAIN_BRANCH
  cd centrifuge-chain
  cargo build --release
  cd ..
fi

./centrifuge-chain/target/release/centrifuge-chain purge-chain --dev
./centrifuge-chain/target/release/centrifuge-chain --dev
