FROM golang:1.13-alpine
RUN apk --no-cache add gcc build-base git make musl-dev linux-headers bash nodejs wget

# Install Go-Ethereum
RUN git clone https://github.com/ethereum/go-ethereum.git geth
RUN cd geth && make devtools

# Install Parity/Subkey
RUN wget -O /usr/bin/Subkey https://storage.googleapis.com/centrifuge-dev-public/subkey
RUN chmod +x /usr/bin/Subkey

RUN ls -la /usr/bin/

RUN /usr/bin/subkey --version

# Copy assets
ADD . /src
WORKDIR /src
RUN make build

CMD /bin/sh -c 'KEYSTORE_PASSWORD=chainsafe ./build/chainbridge'

