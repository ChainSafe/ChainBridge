#!/usr/bin/env bash

# Exit on failure
set -e

PORT=${PORT:-8545}

cd ./on-chain/evm-contracts &&
npm install

echo "Running ganache..."
if [[ $SILENT ]]; then
    node_modules/.bin/ganache-cli -p $PORT -m "$MNEMONIC" --account "0x000000000000000000000000000000000000000000000000000000416c696365,100000000000000000000" --account "0x0000000000000000000000000000000000000000000000000000000000426f62,100000000000000000000" --account "0x00000000000000000000000000000000000000000000000000436861726c6965,100000000000000000000"  --account "0x0000000000000000000000000000000000000000000000000000000044617665,100000000000000000000" --account "0x0000000000000000000000000000000000000000000000000000000000457665,100000000000000000000" > /dev/null &
else
    node_modules/.bin/ganache-cli -p $PORT -m "$MNEMONIC" --account "0x000000000000000000000000000000000000000000000000000000416c696365,100000000000000000000" --account "0x0000000000000000000000000000000000000000000000000000000000426f62,100000000000000000000" --account "0x00000000000000000000000000000000000000000000000000436861726c6965,100000000000000000000"  --account "0x0000000000000000000000000000000000000000000000000000000044617665,100000000000000000000" --account "0x0000000000000000000000000000000000000000000000000000000000457665,100000000000000000000"
fi