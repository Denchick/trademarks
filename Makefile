.PHONY: build
build:
	go build -v ./cmd/xml2csv

.DEFAULT_GOAL := build