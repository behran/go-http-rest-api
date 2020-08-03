.PHONY: build

build:
	go build -v ./src/cmd/apiserver

.PHONY: test

test:
	go test -v -race -timeout 30s ./...


docker.build:
	docker build -t behran/goloang . -f docker/golang/Dockerfile
docker.push:
	docker push -t behran/goloang

.DEFAULT_GOAL := build
