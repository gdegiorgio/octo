all: build test

build:
	go build -o bin/octo cmd/octo/main.go

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out



install:
	go mod tidy

lint:
	golangci-lint run ./...

build-win-amd64:
	set GOOS=windows
	set GOARCH=amd64
	go build -o bin/octo.exe cmd/octo/main.go

run: build
	./bin/octo

test:
	go test ./...
