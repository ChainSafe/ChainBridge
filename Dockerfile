FROM golang:1.13-alpine
RUN apk --no-cache add gcc build-base git make musl-dev linux-headers bash nodejs wget

# Install Go-Ethereum
RUN git clone https://github.com/ethereum/go-ethereum.git geth
RUN cd geth && make devtools

# Install Parity/Subkey
RUN wget -P /usr/bin/ https://storage.googleapis.com/centrifuge-dev-public/subkey
RUN chmod +x /usr/bin/subkey

RUN /usr/bin/subkey --version

ENTRYPOINT ["/usr/bin/subkey"]

# Copy assets
ADD . /src
WORKDIR /src
RUN make build

CMD /bin/sh -c 'KEYSTORE_PASSWORD=chainsafe ./build/chainbridge'

