#!/usr/bin/env bash

pushd ./contracts/evm-contracts &&
npm install &&
node ./scripts/deploy_local.js --validators 3
