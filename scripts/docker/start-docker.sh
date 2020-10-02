#!/bin/bash
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only


set -e

docker build -f ./Docker/Ganache.Dockerfile -t 'chainbridge_ganache:latest' .
docker run -d -p 8545:8545 -p 8546:8546 chainbridge_ganache

# docker exec -it chainbridge_ganache node ./scripts/cli/index.js --relayers 3

dip=`ifconfig | grep 'inet 192'| awk '{ print $2}'`

sed "s/0.0.0.0/$dip/g" config.json > docker.json

cp docker.json config.json

docker build -f ./Docker/Bridge.Dockerfile -t 'chainbridge_bridge:latest' .
docker run chainbridge_bridge
