SHELL := /bin/bash

build/api: # Build api service
	@echo "building api service.."
	go build -o bin/api ./cmd/api/main.go

run/api: # Run api service
	@echo "running api service.."
	go run ./cmd/api

clean: # Remove all build
	rm -fr ./bin