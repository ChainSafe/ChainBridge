FROM node:10

# ENV BASE_URL ganache

RUN mkdir /app
WORKDIR /app

COPY ./contracts/evm-contracts .

RUN yarn install --silent && node_modules/.bin/truffle compile

CMD ["node", "./scripts/deploy_local.js --validators 3"]
