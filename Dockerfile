FROM golang:1.13-alpine as builder

RUN apk --no-cache add gcc build-base git make musl-dev linux-headers bash nodejs

# Install Go-Ethereum
RUN git clone https://github.com/ethereum/go-ethereum.git geth
RUN cd geth && make devtools

ADD . /src
WORKDIR /src
RUN make build

CMD /bin/sh -c 'KEYSTORE_PASSWORD=chainsafe ./build/chainbridge'

