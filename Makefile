.PHONY: run build

run:
	go run go-shortener/cmd/server
build:
	go build -o build/go-shortener-server go-shortener/cmd/server
test:
	go test go-shortener/...
migrate_init:
	goose -dir ./migrations -s postgres "host=localhost user=postgres password=12345678" up

postgres_init:
	docker run --rm --name go-shortener-postgres -p 5432:5432 -e POSTGRES_PASSWORD=12345678 -d postgres

clean:
	goose -dir ./migrations -s postgres "host=localhost user=postgres password=12345678" down
	docker stop go-shortener-postgres