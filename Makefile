SHELL := /bin/bash

build:
	go build -o bin/server ./cmd/server/main.go

run:
	go run ./cmd/server

coverage:
	go test ./... -v -coverprofile=coverage.out && go tool cover -func=coverage.out

clean:
	rm -fr ./bin
	rm coverage.out