all: test build

build:
	go build -o bin/octo cmd/octo/main.go

run: build
	./bin/octo
