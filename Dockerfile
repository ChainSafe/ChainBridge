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

# Add Parity/Subkey
ADD https://storage.googleapis.com/centrifuge-dev-public/subkey /workdir/bridge
RUN cd /workdir/bridge && chmod +x ./subkey && cp subkey /usr/local/bin && subkey --version

ENTRYPOINT cd /workdir/bridge && ls -la && KEYSTORE_PASSWORD=chainsafe ./build/chainbridge