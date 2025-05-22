all: build test

build:
	go build 

test:
	go test

release:
	goreleaser release --clean

.PHONY: all build test
