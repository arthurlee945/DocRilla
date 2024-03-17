build:
	@go build -o bin/docrilla
run:
	@air --build.cmd "go build -o bin/docrilla cmd/main/main.go" --build.bin "./bin/docrilla"
test:
	@go test -v ./...