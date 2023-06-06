.PHONY: run build migrate_init local_postgres_init clean

docker:
	docker compose up
build:
	go build -o build/go-shortener-server go-shortener/cmd/server
test:
	go test go-shortener/...
migrate:
	goose -dir ./migrations up
local_postgres_init:
	docker run --rm --name go-shortener-postgres -p 5432:5432 --env-file .postgres_env -d postgres
generate_proto:
	protoc --go_out=. --go_opt=paths=source_relative \
            --go-grpc_out=. --go-grpc_opt=paths=source_relative \
            api/link/link.proto
clean:
	goose  -dir ./migrations down
	docker stop go-shortener-postgres