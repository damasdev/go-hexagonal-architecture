SHELL := /bin/bash

build/server:
	go build -o bin/server ./cmd/server/main.go

run/server:
	go run ./cmd/server

build/consumer:
	go build -o bin/consumer ./cmd/consumer/main.go

run/consumer:
	go run ./cmd/consumer

coverage:
	go test ./... -v -coverprofile=coverage.out && go tool cover -func=coverage.out

clean:
	rm -fr ./bin
	rm coverage.out