#!/usr/bin/env bash

pushd ./on-chain/evm-contracts &&
npm install &&
node ./scripts/deploy_local.js --validators 3
