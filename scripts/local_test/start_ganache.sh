#!/usr/bin/env bash

# Exit on failure
set -e

PORT=${PORT:-8545}

cd ./contracts/evm-contracts &&
npm install

echo "Running ganache..."
if [[ $1 = "silent" ]]; then
    node_modules/.bin/ganache-cli -h 0.0.0.0 -p $PORT -m "when sound uniform light fee face forum huge impact talent exhaust arrow" > /dev/null &
else
    node_modules/.bin/ganache-cli -h 0.0.0.0 -p $PORT -m "when sound uniform light fee face forum huge impact talent exhaust arrow"
fi
