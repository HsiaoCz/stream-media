run:build
	@go run ./bin/stream-media

build:
	@go build -o bin/stream-media

test:
	@go test -v ./...

all:run

.PHONY:run build test all

help:
	@echo "run :go run ./bin/stream-dedia"
	@echo "build :go build bin/stream-dedia"
	@echo "test :go test ./..."
	