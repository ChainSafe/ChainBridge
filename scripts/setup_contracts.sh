#!/usr/bin/env bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only

CONTRACTS_REPO="https://github.com/ChainSafe/chainbridge-solidity"
CONTRACTS_BRANCH="master"
CONTRACTS_COMMIT="1bfd1fba73c9c93a828b680577683c79429a56eb"
CONTRACTS_DIR="./solidity"
DEST_DIR="./bindings"

set -e

case $TARGET in
	"build")
		git clone -b $CONTRACTS_BRANCH $CONTRACTS_REPO $CONTRACTS_DIR
    pushd $CONTRACTS_DIR
    git checkout $CONTRACTS_COMMIT

    make install-deps
    make bindings

    popd

    mkdir $DEST_DIR
    cp -r $CONTRACTS_DIR/build/bindings/go/* $DEST_DIR
		;;

	"cli-only")
		git clone -b $CONTRACTS_BRANCH $CONTRACTS_REPO $CONTRACTS_DIR
    pushd $CONTRACTS_DIR
    git checkout $CONTRACTS_COMMIT
		;;


esac

