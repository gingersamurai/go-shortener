.PHONY: run build migrate_init local_postgres_init clean

docker:
	docker compose up

run:
	go run go-shortener/cmd/server
build:
	go build -o build/go-shortener-server go-shortener/cmd/server
test:
	go test go-shortener/...
migrate:
	goose -dir ./migrations up

local_postgres_init:
	docker run --rm --name go-shortener-postgres -p 5432:5432 --env-file .postgres_env -d postgres

clean:
	goose  -dir ./migrations down
	docker stop go-shortener-postgres