FROM golang:latest

ARG PATH_BIN=/usr/local/bin

RUN set -xe \
    && apt-get update && apt-get install -y \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

COPY ./src/ /go/src/app/

WORKDIR /go/src/app/

RUN go mod download \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app cmd/apiserver/main.go \
    && cp app $PATH_BIN/app \
    && chmod 755 $PATH_BIN/app

ENTRYPOINT ["/usr/local/bin/app"]

EXPOSE 80
