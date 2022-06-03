FOLDER_MIGRATIONS = src/migrations
FOLDER_SERVER = src/server

gen-error-code:
	go run src/server/main.go gen-error-code

install-tools:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

migrate-create:
	@go run $(FOLDER_SERVER)/main.go migrate-create ${name}

migrate:
	@go run $(FOLDER_SERVER)/main.go migrate

seed:
	@go run $(FOLDER_SERVER)/main.go seed-${name}

gen-post:
	protoc src/proto/post/post.proto --go_out=paths=source_relative:./ --go-grpc_out=paths=source_relative:./

gen-class:
	protoc src/proto/class/class.proto --go_out=paths=source_relative:./ --go-grpc_out=paths=source_relative:./