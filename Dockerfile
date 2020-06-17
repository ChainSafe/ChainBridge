# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only
FROM  golang:1.13-alpine AS builder
RUN apk --no-cache add gcc build-base git linux-headers
ADD . /src
WORKDIR /src
RUN go mod download
RUN cd cmd/chainbridge && go build -o /bridge .

## Final stage
## Start with subkey build
FROM parity/subkey:2.0.0-alpha.3
USER root
RUN subkey --version
COPY --from=builder /bridge ./
RUN chmod +x ./bridge

ENTRYPOINT ["./bridge"]
