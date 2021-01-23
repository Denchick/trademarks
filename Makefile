.PHONY: build
build:
	go build -v ./cmd/initdb

.DEFAULT_GOAL := build