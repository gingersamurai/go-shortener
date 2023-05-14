package main

import (
	"go-shortener/internal/delivery/http_server"
	"go-shortener/internal/storage/postgres_storage"
	"go-shortener/internal/usecase"
	"go-shortener/internal/usecase/shortener"
	"go-shortener/pkg/closer"
	"log"
	"time"
)

const (
	host            = "localhost:8080"
	handleTimeout   = time.Second * 5
	shutdownTimeout = time.Second * 2
)

func main() {

	appCloser := closer.NewCloser(shutdownTimeout)

	appStorage, err := postgres_storage.NewPostgresStorage("host=localhost user=postgres password=12345678")
	if err != nil {
		log.Fatal(err)
	}
	appCloser.Add(appStorage.Shutdown)

	//memoryStorage := memory_storage.NewMemoryStorage()

	appShortener, err := shortener.NewPolynomialHashShortener(10)
	if err != nil {
		log.Fatal(err)
	}

	linkInteractor := usecase.NewLinkInteractor(appShortener, appStorage)

	handler := http_server.NewHandler(host, handleTimeout, linkInteractor)
	server := http_server.NewServer(host, handler)
	appCloser.Add(server.Shutdown)

	go server.Run()

	appCloser.Run()

}
