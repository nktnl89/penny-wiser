.PHONY: build
build: go build -v ./penny-wiser/apiserver

.DEFAULT_GOAL := build