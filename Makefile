
run: move_migrations
	@go run .

deps:
	@go mod tidy

build: move_migrations
	@go build -o memo

test:
	@go test -v

querie:
	@sqlc generate

move_migrations:
	@cp -r db/migrations ~/.memo/
