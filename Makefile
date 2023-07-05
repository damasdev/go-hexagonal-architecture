SHELL := /bin/bash

build:
	go build -o bin/api ./cmd/api/main.go

run:
	go run ./cmd/api

compile:
	GOOS=linux GOARCH=arm go build -o bin/api-linux-arm ./cmd/api/main.go
	GOOS=linux GOARCH=arm64 go build -o bin/api-linux-arm64 ./cmd/api/main.go