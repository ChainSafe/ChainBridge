#!/usr/bin/env bash

cd ./contracts/evm-contracts &&
npm install

echo "Running ganache..."
if [ $1 = "silent" ]; then
    node_modules/.bin/ganache-cli -m when sound uniform light fee face forum huge impact talent exhaust arrow > /dev/null &
else 
    node_modules/.bin/ganache-cli -m when sound uniform light fee face forum huge impact talent exhaust arrow
fi
