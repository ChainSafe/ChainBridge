# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only

FROM node:12.14-alpine3.11

RUN mkdir /app
WORKDIR /app

COPY ./on-chain/evm-contracts .

RUN yarn install --silent && node_modules/.bin/truffle compile

ENTRYPOINT ["node_modules/.bin/ganache-cli", "-h", "0.0.0.0", "-p","8545","-m","when sound uniform light fee face forum huge impact talent exhaust arrow"]

