-include .env
export

run:
	go run main.go

fmt:
	go fmt ./...

lint:
	golangci-lint run -c .golangci.yml

dep:
	go mod download

coverage:
	go test -cover -coverprofile=coverage.out ./... && \
	 go tool cover -html=coverage.out
