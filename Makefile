
run:
	@go run .

deps:
	@go mod tidy

build:
	@go build -o memo

test:
	@go test -v
