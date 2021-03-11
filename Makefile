-include .env
export

test: dep
	go test ./...

run:
	go run main.go

fmt:
	go fmt ./...

lint:
	golangci-lint run -c .golangci.yml

dep:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.38.0 && \
	go mod download

coverage:
	go test -cover -coverprofile=coverage.out ./... && \
	 go tool cover -html=coverage.out
