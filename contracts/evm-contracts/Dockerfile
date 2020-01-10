FROM node:10

ENV BASE_URL ganache

RUN mkdir /app
WORKDIR /app

COPY . .

RUN yarn install --silent && node_modules/.bin/truffle compile

CMD ["sleep", "infinity"]

# CMD ["node", "./scripts/deploy_local.js --validators 3"]
