SHELL := /bin/bash

build/api:
	go build -o bin/api ./cmd/api/main.go

run/api:
	go run ./cmd/api

coverage:
	go test ./... -v -coverprofile=coverage.out && go tool cover -func=coverage.out

clean:
	rm -fr ./bin
	rm coverage.out