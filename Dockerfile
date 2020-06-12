# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only
FROM  golang:1.13-alpine AS builder
RUN apk --no-cache add gcc build-base git linux-headers
ADD . /src
WORKDIR /src
RUN go mod download
RUN cd cmd/chainbridge && go build -o /bridge .

# # final stage
FROM debian:stretch-slim
RUN apt-get -y update && apt-get -y upgrade && apt-get install ca-certificates wget -y
RUN wget -P /usr/local/bin/ https://storage.googleapis.com/centrifuge-dev-public/subkey  && chmod +x /usr/local/bin/subkey
RUN subkey --version

COPY --from=builder /bridge ./
RUN chmod +x ./bridge

ENV KEYSTORE_PASSWORD=chainsafe

ENTRYPOINT ["./bridge"]
