build:
	@go build -o cmd/builds
run: build
	@./cmd/builds
	@DB_MIGRATE=true go run main.go
test:
	@go test -v ./...
    
