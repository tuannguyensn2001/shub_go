FOLDER_MIGRATIONS = src/migrations
FOLDER_SERVER = src/server

gen-error-code:
	go run src/server/main.go gen-error-code

install-tools:
	go install github.com/pressly/goose/v3/cmd/goose@latest

migrate-create:
	@go run $(FOLDER_SERVER)/main.go migrate-create ${name}

migrate:
	@go run $(FOLDER_SERVER)/main.go migrate

seed:
	@go run $(FOLDER_SERVER)/main.go seed-${name}

migrate-down:
	@go run $(FOLDER_SERVER)/main.go migrate-down