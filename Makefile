
run:
	@go run .

deps:
	@go mod tidy

build:
	@go build -o memo

test:
	@go test -v

querie:
	@sqlc generate

migrate:
	@goose -dir=db/migrations sqlite3 test.db up

migrate_status:
	@goose -dir=db/migrations sqlite3 test.db status

migrate_revert:
	@goose -dir=db/migrations sqlite3 test.db down
