.PHONY: test build

test:
	go test -v ./...

run:
	go run cmd/cli/main.go

build:
	go build -o bin/unit-converter cmd/cli/main.go
