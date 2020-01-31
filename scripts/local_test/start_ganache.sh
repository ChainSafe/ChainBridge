#!/usr/bin/env bash

# Exit on failure
set -e

PORT=${PORT:-8545}

cd ./on-chain/evm-contracts &&
npm install

echo "Running ganache..."
if [[ $1 = "silent" ]]; then
    node_modules/.bin/ganache-cli -p $PORT -m "when sound uniform light fee face forum huge impact talent exhaust arrow" > /dev/null &
else
    node_modules/.bin/ganache-cli -p $PORT -m "when sound uniform light fee face forum huge impact talent exhaust arrow" --account "0x000000000000000000000000000000000000000000000000000000416c696365,0x56bc75e2d63100000" "0x0000000000000000000000000000000000000000000000000000000000426f62
,0x56bc75e2d63100000"
fi
