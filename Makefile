.PHONY: test build

test:
	go test -v ./...

run:
	go run cmd/unit-converter/main.go

build:
	go build -o bin/unit-converter ./cmd/unit-converter

install:
	go install ./cmd/unit-converter
