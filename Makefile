WORKDIR = /home/gingersamurai/coding/projects/go-shortener

.PHONY: run build migrate_init local_postgres_init clean

build:
	go build -o $(WORKDIR)/build/go-shortener-server $(WORKDIR)/cmd/server

# TESTS
lint_test:
	golangci-lint run $(WORKDIR)/...
unit_test:
	go test $(WORKDIR)/...

migrate:
	goose -dir $(WORKDIR)/migrations up

#LOCAL DEVELOPMENT
local_postgres_init:
	docker run --rm --name go-shortener-postgres -p 5432:5432 --env-file .postgres_env -d postgres
generate_proto:
	protoc --go_out=. --go_opt=paths=source_relative \
            --go-grpc_out=. --go-grpc_opt=paths=source_relative \
            api/link/link.proto

clean:
	goose  -dir ./migrations down
	docker stop go-shortener-postgres
	rm -rf