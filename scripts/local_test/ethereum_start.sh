#!/usr/bin/env bash

pushd ./contracts/evm-contracts &&
npm install &&
# TODO Can conifigure number of validators
# node ./scripts/deploy_local.js 4
node ./scripts/deploy_local.js
