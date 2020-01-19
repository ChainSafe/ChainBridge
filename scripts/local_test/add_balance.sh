#!/usr/bin/env bash

dip=`ifconfig | grep 'inet 192'| awk '{ print $2}'`
ADMIN_ADDR=${ADMIN_ADDR:-'0x34c59fBf82C9e31BA9CBB5faF4fe6df05de18Ad4'}
TO_ADDR=${TO_ADDR:-'0x56e77cD98C241b0Fb70Bc83A8eF41D94a30C6f1e'}
docker run ethereum/client-go:latest attach ws://${dip}:8545 --exec "personal.unlockAccount('$ADMIN_ADDR', '', 500)"
docker run ethereum/client-go:latest attach ws://${dip}:8545 --exec "eth.sendTransaction({from:'${ADMIN_ADDR}', to:'${TO_ADDR}', value: web3.toWei(10, 'ether'), gas:21000});"
