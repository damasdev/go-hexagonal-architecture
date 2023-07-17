SHELL := /bin/bash

build:
	go build -o bin/api ./cmd/api/main.go

run:
	go run ./cmd/api

coverage:
	go test ./... -v -coverprofile=coverage.out && go tool cover -func=coverage.out

clean:
	rm -fr ./bin
	rm coverage.out