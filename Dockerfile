FROM golang:1.18-bookworm AS builder

ARG TARGETARCH
ARG flydb_version=1.0.6

ENV FLYDB_VERSION=${flydb_version} \
    FLYDB_HOME=/flydb

WORKDIR /src

RUN git clone -b "v${FLYDB_VERSION}" https://github.com/ByteStorage/FlyDB.git && \
    cd FlyDB
ENV GOPROXY=https://proxy.golang.org
WORKDIR /src/FlyDB

RUN GO111MODULE=on go mod download
RUN go build -o /flydb/flydb-server -ldflags="-s -w" cmd/server/cli/flydb-server.go

FROM debian:bookworm-slim

RUN apt-get update &&  \
    apt-get upgrade -y &&  \
    apt-get install tzdata

COPY --from=builder /flydb/flydb-server /flydb/flydb-server
WORKDIR /flydb

CMD [ "/flydb/flydb-server" ]

EXPOSE 8999