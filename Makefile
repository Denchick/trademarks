.PHONY: build
build:
	go build -v ./cmd/xml2csv
	go build -v ./cmd/server

.DEFAULT_GOAL := build