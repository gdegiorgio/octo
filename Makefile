all: test build

build:
	go build -o bin/octo cmd/octo/main.go

build-win-amd64:
	set GOOS=windows
	set GOARCH=amd64
	go build -o bin/octo.exe cmd/octo/main.go

run: build
	./bin/octo

test:
	go test ./...
