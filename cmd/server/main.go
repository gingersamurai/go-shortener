package main

import (
	"go-shortener/internal/delivery/http_server"
	"go-shortener/internal/storage/postgres_storage"
	"go-shortener/internal/usecase"
	"go-shortener/internal/usecase/shortener"
	"log"
	"time"
)

func main() {
	host := "localhost:8080"
	handleDeadline := time.Second * 15

	postgresStorage, err := postgres_storage.NewPostgresStorage("host=localhost user=postgres password=12345678")
	if err != nil {
		log.Fatal(err)
	}

	polynomialHashShortener, err := shortener.NewPolynomialHashShortener(10)
	if err != nil {
		log.Fatal(err)
	}

	linkInteractor := usecase.NewLinkInteractor(polynomialHashShortener, postgresStorage)

	handler := http_server.NewHandler(host, handleDeadline, linkInteractor)
	server := http_server.NewServer(host, handler)

	log.Fatal(server.Run())
}
