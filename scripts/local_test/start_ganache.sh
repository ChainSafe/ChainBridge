#!/usr/bin/env bash

# Exit on failure
set -e

pushd ./contracts/evm-contracts &&
npm install &&
node_modules/.bin/ganache-cli -m when sound uniform light fee face forum huge impact talent exhaust arrow