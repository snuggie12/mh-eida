build: tidy
	go build -o bin/

tidy:
	go mod tidy

all: build
.PHONY: build tidy
.DEFAULT_GOAL := all
