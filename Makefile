build:
	@go build -o tmp cmd/server/main.go
run:
	@air --build.cmd "make build"
test:
	@go test -v ./internal/...