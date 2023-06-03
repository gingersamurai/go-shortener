package main

import (
	"errors"
	"go-shortener/internal/config"
	"go-shortener/internal/delivery/grpc_server"
	"go-shortener/internal/delivery/http_server"
	"go-shortener/internal/storage/memory_storage"
	"go-shortener/internal/storage/postgres_storage"
	"go-shortener/internal/usecase"
	"go-shortener/pkg/closer"
	"go-shortener/pkg/polynomial_hash_shortener"
	"log"
)

func chooseStorage(appConfig config.Config, appCloser *closer.Closer) (usecase.Storage, error) {
	switch appConfig.StorageType {
	case "memory":
		return memory_storage.NewMemoryStorage(), nil
	case "postgres":
		postgresStorage, err := postgres_storage.NewPostgresStorage(appConfig.Postgres)
		if err != nil {
			return nil, err
		}
		appCloser.Add(postgresStorage.Shutdown)
		return postgresStorage, nil
	default:
		return nil, errors.New("wrong storage type")
	}
}

func main() {
	appConfig, err := config.NewConfig(config.ConfigFilePath, config.ConfigFileName)
	if err != nil {
		log.Fatal(err)
	}
	appCloser := closer.NewCloser(appConfig.ShutdownTimeout)

	appStorage, err := chooseStorage(appConfig, appCloser)
	if err != nil {
		log.Fatal(err)
	}

	appShortener, err := polynomial_hash_shortener.NewPolynomialHashShortener(10)
	if err != nil {
		log.Fatal(err)
	}

	linkInteractor := usecase.NewLinkInteractor(appShortener, appStorage)

	handler := http_server.NewHandler(appConfig.Handler, linkInteractor)
	server := http_server.NewServer(appConfig.HttpServer, handler)
	appCloser.Add(server.Shutdown)
	go server.Run()

	grpcHandler := grpc_server.NewHandler(appConfig.Handler, linkInteractor)
	grpcServer, err := grpc_server.NewServer(appConfig.GrpcServer, grpcHandler)
	if err != nil {
		log.Fatal(err)
	}
	appCloser.Add(grpcServer.Shutdown)
	go grpcServer.Run()
	log.Println("started grpc with separate goroutine")
	appCloser.Run()

}
