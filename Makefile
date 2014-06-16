.PHONY: deps test build

BINARY := battlestations

all: deps test build

deps:
		go get -t -v ./...

build:
		go build -o $(BINARY)

test:
		go test -v ./...
