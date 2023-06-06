package main

import (
	"errors"
	"flag"
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

func getConfigPath() (string, error) {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()
	if configPath == nil {
		return "", errors.New("bad config path")
	}
	return *configPath, nil
}

func main() {
	configPath, err := getConfigPath()
	if err != nil {
		log.Fatal(err)
	}
	appConfig, err := config.NewConfig(configPath)
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

	appCloser.Run()

}
