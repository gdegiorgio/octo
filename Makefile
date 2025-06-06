all: build test

build:
	goreleaser build --snapshot --clean

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

install:
	go mod tidy

lint:
	golangci-lint run ./...

run: build
	./bin/octo

test:
	go test ./...
