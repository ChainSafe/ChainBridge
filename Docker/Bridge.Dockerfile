FROM  golang:1.12.5-alpine AS builder
RUN apk --no-cache add gcc build-base git
ADD . /bridge
WORKDIR /bridge
RUN go mod download
RUN cd cmd/chainbridge && go build -o ../../bridge
RUN ls -la

# final stage
FROM alpine:3.11
COPY --from=builder /bridge /
# ENTRYPOINT [ "bridge" ]