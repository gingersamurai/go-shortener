package main

import (
	"go-shortener/internal/delivery/http_server"
	"go-shortener/internal/storage/memory_storage"
	"go-shortener/internal/usecase"
	"go-shortener/internal/usecase/shortener"
	"log"
)

func main() {
	host := "localhost:8080"
	memoryStorage := memory_storage.NewMemoryStorage()

	polynomialHashShortener, err := shortener.NewPolynomialHashShortener(10)
	if err != nil {
		log.Fatal(err)
	}

	linkInteractor := usecase.NewLinkInteractor(polynomialHashShortener, memoryStorage)

	handler := http_server.NewHandler(host, linkInteractor)
	server := http_server.NewServer(host, handler)

	log.Fatal(server.Run())
}
