FROM  golang:1.12.5-alpine AS builder
RUN apk --no-cache add gcc build-base git
ADD . /src
WORKDIR /src
RUN go mod download
RUN cd cmd/chainbridge && go build -o /bridge .

# # final stage
FROM alpine:latest
# RUN apk --no-cache add ca-certificates curl
COPY --from=builder /bridge /src/config.toml /src/scripts/docker/wait-for-ganache.sh ./
COPY --from=builder /src/keys/ ./keys/
RUN chmod +x ./bridge ./keys/ ./wait-for-ganache.sh
