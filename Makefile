.PHONY:
build:
	go build -v ./cmd

.PHONY:
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build
