build:
	@go build -o cmd/builds
run: build
	@./cmd/builds
test:
	@go test -v ./...