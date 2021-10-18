SHELL=/bin/bash
.PHONY: all help run build test version prerelease patch minor major clean

all: build

help:
	go run main.go -h

run:
	go run main.go -port 8080 -verbosity vvvvvv

version:
	go run main.go -version

build:
	go build -o avalon
	tar czf avalon.tar.gz avalon html server

test:
	go test github.com/damontic/avalon/tests/server
	go test github.com/damontic/avalon/tests/server/handlers

clean:
	rm avalon.tar.gz avalon
