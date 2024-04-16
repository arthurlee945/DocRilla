build:
	@go build -o bin/docrilla cmd/server/main.go
run:
	@air --build.cmd "make build" --build.bin "./bin/docrilla"
test:
	@go test -v ./internal/...