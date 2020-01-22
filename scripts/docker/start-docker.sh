#!/bin/bash

set -e

docker build -f ./Docker/Ganache.Dockerfile -t 'chainbridgev2_ganache:latest' .
docker run -d -p 8545:8545 -p 8546:8546 chainbridgev2_ganache

# docker exec -it chainbridgev2_ganache node ./scripts/deploy_local.js --validators 3

dip=`ifconfig | grep 'inet 192'| awk '{ print $2}'`

sed "s/0.0.0.0/$dip/g" config.toml > docker.toml

cp docker.toml config.toml

docker build -f ./Docker/Bridge.Dockerfile -t 'chainbridgev2_bridge:latest' .
docker run chainbridgev2_bridge
