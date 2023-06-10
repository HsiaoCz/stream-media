run:build
	@./bin/stream-media

build:
	@go build -o bin/stream-media

test:
	@go test -v ./...

help:
	@echo "run :./bin/stream-dedia"
	@echo "build :go build bin/stream-dedia"
	@echo "test :go test ./..."

all: run

.PHONY:run build test all help