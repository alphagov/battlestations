.PHONY: deps keys test build

BINARY := battlestations

all: deps test build

deps:
		go get -t -v ./...

keys:
		cat /dev/urandom | head -c 64 > auth.key
		cat /dev/urandom | head -c 32 > enc.key

build:
		go build -o $(BINARY)

test:
		go test -v ./...
