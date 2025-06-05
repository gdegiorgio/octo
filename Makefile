all: test build

build:
	go build -o bin/octo cmd/main.go

run: build
	./bin/octo
