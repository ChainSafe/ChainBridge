#!/usr/bin/env bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only

DIR="./build/centrifuge-chain"
CENT_CHAIN_URL="https://github.com/ansermino/centrifuge-chain.git"
CENT_CHAIN_BRANCH="david/bridge-module"

# Exit on failure
set -e

if [ "$CMD" == "fetch" ]
then
  if [ ! -d $DIR ] || [ -z "$(ls -A $DIR)" ]
  then
      git clone $CENT_CHAIN_URL -b $CENT_CHAIN_BRANCH $DIR
  else
    cd $DIR
    git pull origin $CENT_CHAIN_BRANCH
  fi
elif [ "$CMD" == "build" ]
then
  cd $DIR
  RUST_TOOLCHAIN=nightly-2020-02-17 TARGET=build-client ./ci/script.sh
elif [ "$CMD" == "run" ]
then
  ls -la $DIR
  ls -la $DIR/target/release
  $DIR/target/release/centrifuge-chain --dev > /dev/null
else
  echo "Please provide CMD={build,run}"
fi
