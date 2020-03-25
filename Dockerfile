// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

FROM golang:1.13-buster

WORKDIR /workdir

# Node 12 setup
RUN curl -sL https://deb.nodesource.com/setup_12.x -o nodesource_setup.sh
RUN chmod +x ./nodesource_setup.sh && ./nodesource_setup.sh && apt install -y nodejs
RUN node -v && npm -v

# Install Go-Ethereum
RUN git clone https://github.com/ethereum/go-ethereum.git geth
RUN cd geth && make devtools

# Copy bridge
ADD . /workdir/bridge
RUN cd /workdir/bridge && make build

# Rename default config
RUN ls -la
RUN mv bridge/config.toml.example config.toml
RUN mv bridge/keys keys

# Add Parity/Subkey
ADD https://storage.googleapis.com/centrifuge-dev-public/subkey /workdir/bridge
RUN cd /workdir/bridge && chmod +x ./subkey && cp subkey /usr/local/bin && subkey --version

ENTRYPOINT KEYSTORE_PASSWORD=chainsafe ./bridge/build/chainbridge
