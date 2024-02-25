.PHONY: build run test docker-build docker-run

build:
	go build -o plutusapi ./cmd/api/main.go 

run: build
	./plutusapi 

test:
	go test -race ./...

docker-build:
	docker build -t plutusapi .

docker-run: docker-build
	docker run -p 8080:8080 plutusapi
