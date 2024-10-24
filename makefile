.PHONY: all test clean

test:
	go test .internal/server/ -v

lint:
	golangci-lint run

run:
	go run ./cmd/movietracker/main.go

build:
	docker build --pull --rm -f "Dockerfile" -t movietracker:latest "."

up:
	docker run -p 8080:8080 movietracker:latest