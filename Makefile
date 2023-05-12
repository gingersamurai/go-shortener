.PHONY: run build

run:
	go run go-shortener/cmd/server
build:
	go build -o build/go-shortener-server go-shortener/cmd/server
test:
	go test go-shortener/...