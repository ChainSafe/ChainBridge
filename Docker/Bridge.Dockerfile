FROM  golang:1.13-alpine AS builder
RUN apk --no-cache add gcc build-base git
ADD . /src
WORKDIR /src
RUN go mod download
RUN cd cmd/chainbridge && go build -o /bridge .

# # final stage
FROM alpine:latest
# RUN apk --no-cache add ca-certificates curl
COPY --from=builder /bridge /src/config.toml ./
COPY --from=builder /src/keys/ ./keys/
RUN chmod +x ./bridge ./keys/
CMD /bin/sh -c 'KEYSTORE_PASSWORD=chainsafe ./bridge'