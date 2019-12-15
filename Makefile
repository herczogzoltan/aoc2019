.PHONY: build get-deps update-deps test help default

default: help

BIN_NAME="aoc2019"

VERSION := $(shell grep "const Version " version/version.go | sed -E 's/.*"(.+)".$$/\1/')
GIT_COMMIT = $(shell git rev-parse HEAD)
BUILD_DATE = $(shell date '+%Y-%m-%d-%H:%M:%S')
REPOSITORY = github.com/herczogzoltan/aoc2019

help:
	@echo 'Directives for AOC:'
	@echo
	@echo 'Usage:'
	@echo '    make build           Compile the project.'
	@echo '    make get-deps        runs mod download to download all the dependencies'
	@echo '    make update-deps     runs mod tidy to update dependencies'
	@echo '    make test            Run tests on a compiled project.'
	@echo

build:
	@echo "building ${BIN_NAME} ${VERSION}"
	go build -ldflags "-X ${REPOSITORY}/version.GitCommit=${GIT_COMMIT} -X ${REPOSITORY}/version.BuildDate=${BUILD_DATE}" -o bin/${BIN_NAME}
	
get-deps:
	@echo "Downloading dependencies..."
	go mod download

update-deps:
	@echo "Updating dependencies..."
	go mod tidy

test:
	go test ./... -v