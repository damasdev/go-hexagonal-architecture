SHELL := /bin/bash

build/api:
	go build -o bin/api ./cmd/api/main.go

run/api:
	go run ./cmd/api

test:
	go test -v ./test/..

coverage:
	go test ./test/... -coverprofile=cover.out -coverpkg ./... && go tool cover -func=cover.out

clean:
	rm -fr ./bin