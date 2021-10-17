SHELL=/bin/bash
.PHONY: all run build test version prerelease patch minor major clean

all: publish

help:
	go run main.go -h

run:
	go run main.go -port 8080

build:
	go build -o avalon
	tar czf avalon.tar.gz avalon html server

test:
	go test github.com/damontic/avalon/tests/server
	go test github.com/damontic/avalon/tests/server/handlers

version:
	@cat version

prerelease: test
	@bash -c 'semver -i prerelease --preid beta $$(cat version) > version'

patch: test
	@bash -c 'semver -i patch $$(cat version) > version'

minor: test
	@bash -c 'semver -i minor $$(cat version) > version'

major: test
	@bash -c 'semver -i major $$(cat version) > version'

clean:
	rm avalon.tar.gz avalon
