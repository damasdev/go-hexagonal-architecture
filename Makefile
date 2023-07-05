SHELL := /bin/bash

build/api:
	@echo "building api service.."
	go build -o bin/api ./cmd/api/main.go

run/api:
	@echo "running api service.."
	go run ./cmd/api