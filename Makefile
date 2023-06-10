run:build
	@./bin/stream-media

build:
	@go build -o bin/stream-media

test:
	@go test -v ./...

init:
	@go mod tidy

help:
	@echo "run :./bin/stream-dedia"
	@echo "build :go build bin/stream-dedia"
	@echo "test :go test ./..."
	@echo "init : go mod tidy"

all: run

.PHONY:run build test all help