# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only

FROM golang:1.13-buster

WORKDIR /workdir

# Docker doesn't accept runtime vairables, so we store the variable at image buildtime
ARG PASSWORD 
ENV KEYSTORE_PSWD=$PASSWORD

# Allows a user to pass in a key from the keyring eg: --alice
ARG TEST_KEY
ENV TEST_KEY_CMD=$TEST_KEY

# Node 12 setup
RUN curl -sL https://deb.nodesource.com/setup_12.x -o nodesource_setup.sh
RUN chmod +x ./nodesource_setup.sh && ./nodesource_setup.sh && apt install -y nodejs
RUN node -v && npm -v

# Install Go-Ethereum
RUN git clone https://github.com/ethereum/go-ethereum.git geth
RUN cd geth && make devtools

# Copy bridge
ADD . /workdir/bridge
RUN cd /workdir/bridge && make setup-contracts && make build

# Rename default config
RUN ls -la
RUN mv bridge/config.toml.example config.toml
RUN mkdir -p ./keys/; mv bridge/keys/* $_

# Add Parity/Subkey
ADD https://storage.googleapis.com/centrifuge-dev-public/subkey /workdir/bridge
RUN cd /workdir/bridge && chmod +x ./subkey && cp subkey /usr/local/bin && subkey --version

ENTRYPOINT KEYSTORE_PASSWORD=$KEYSTORE_PSWD ./bridge/build/chainbridge $TEST_KEY_CMD
