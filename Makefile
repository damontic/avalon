SHELL=/bin/bash
.PHONY: all help run version build test clean

all: build

help:
	bazel run //:avalon -- -h

run:
	bazel run //:avalon -- -port 8080 -verbosity vvvvvv

version:
	bazel run //:avalon -- -version

build:
	bazel build //:avalon_package

test:
	bazel test //tests/server:server_test //tests/server/handlers:handlers_test

clean:
	bazel clean --expunge
