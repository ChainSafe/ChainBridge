#!/bin/bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only


set -e

ARG=$(echo $1 | sed 's/=.*//')
VALUE=$(echo $1 | sed 's/.*=//')

PSWD=""
DEV_KEY=""

if [ -z $ARG ];then
    echo "
    USAGE
    ./start-docker.sh [...options]

    Flags
    --password   Password for the keystore
    --testkey    Test key to use for chainbridge (alice, bob, ...etc)

    Note: Must use one of --password or --testkey
    "
elif [ $ARG == "--password" ];then
    PSWD=$VALUE
elif [ $ARG == "--testkey" ];then
    DEV_KEY=$VALUE
fi

docker image build -t chainbridge . --build-arg PASSWORD=$PSWD --build-arg TEST_KEY="$DEV_KEY"
docker run chainbridge -d -v ./keys:/workdir
