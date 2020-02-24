#!/usr/bin/env bash

pushd ./ethereum &&
npm install &&
node ./scripts/deploy_local.js --validators 3
