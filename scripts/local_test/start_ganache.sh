#!/usr/bin/env bash

# Exit on failure
set -e

PORT=${PORT:-8545}

MNEMONIC="\"when sound uniform light fee face forum huge impact talent exhaust arrow\""

cd ./on-chain/evm-contracts &&
npm install

echo "Running ganache..."
if [[ $SILENT ]]; then
    node_modules/.bin/ganache-cli -p $PORT -m "when sound uniform light fee face forum huge impact talent exhaust arrow" > /dev/null &
else
    node_modules/.bin/ganache-cli -p $PORT -m "when sound uniform light fee face forum huge impact talent exhaust arrow"
fi
