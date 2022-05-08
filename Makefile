.PHONY: build

build:
	go build -v ./cmd/ssr

.DEFAULT_GOAL := build
