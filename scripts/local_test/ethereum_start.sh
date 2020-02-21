#!/usr/bin/env bash

pushd ./ethereum/evm-contracts &&
npm install &&
node ./scripts/deploy_local.js --validators 3
