#!/bin/bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only


set -e

PSWD=$1

docker build -f ./Dockerfile --build-arg PASSWORD=$PSWD -t 'chainbridgev2_ganache:latest' .
docker run -d -v ./keys:/workdir chainbridgev2_ganache
