all: build test

build:
	go build 

test:
	go test

.PHONY: all build test
