FROM golang:1.20-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/evoaway/erc20-transfers-storage-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/erc20-transfers-storage-svc /go/src/github.com/evoaway/erc20-transfers-storage-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/erc20-transfers-storage-svc /usr/local/bin/erc20-transfers-storage-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["erc20-transfers-storage-svc"]
