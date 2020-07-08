#!/usr/bin/env sh
# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only

# Exit on failure
set -ex
geth init --datadir /root/ /root/genesis.json

geth --datadir /root/ --maxpeers 0 --light.maxpeers 0 account import /root/key.txt
#rm -rf /root/.ethereum/keystore
#cp -r /root/keystore /root/.ethereum/

echo groot

exec geth --datadir \
  /root/ \
  --maxpeers 0 \
  --light.maxpeers 0 \
  --unlock 0 \
  --mine console \
  --password /root/password.txt


#exec geth \
#    --nodiscover \
#    --unlock "0xff93B45308FD417dF303D6515aB04D9e89a750Ca","0x8e0a907331554AF72563Bd8D43051C2E64Be5d35","0x24962717f8fA5BA3b931bACaF9ac03924EB475a0","0x148FfB2074A9e59eD58142822b3eB3fcBffb0cd7","0x4CEEf6139f00F9F4535Ad19640Ff7A0137708485" \
#    --password /root/password.txt \
#    --ws \
#    --wsport 8546 \
#    --wsorigins="*" \
#    --wsaddr 0.0.0.0 \
#    --rpc \
#    --rpcport 8545 \
#    --rpccorsdomain="*" \
#    --rpcaddr 0.0.0.0 \
#    --networkid 5 \
#    --targetgaslimit 8000000 \
#    --allow-insecure-unlock \
#    --mine \
#    --keystore /root/keystore \
