install:
	@go get

build:
	@go build -o cmd/streamy-api

cache:
	@rm -f ./cmd/streamy-api

run: cache build
	@./cmd/streamy-api
	@DB_MIGRATE=true go run main.go

test:
	@go test -v ./...

tidy:
	@go mod tidy

format:
	@gofmt -s -w .
