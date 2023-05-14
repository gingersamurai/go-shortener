package main

import (
	"errors"
	"github.com/davecgh/go-spew/spew"
	"go-shortener/internal/config"
	"go-shortener/internal/delivery/http_server"
	"go-shortener/internal/storage/memory_storage"
	"go-shortener/internal/storage/postgres_storage"
	"go-shortener/internal/usecase"
	"go-shortener/internal/usecase/shortener"
	"go-shortener/pkg/closer"
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
	spew.Dump(appConfig)
	appCloser := closer.NewCloser(appConfig.ShutdownTimeout)

	appStorage, err := chooseStorage(appConfig, appCloser)
	if err != nil {
		log.Fatal(err)
	}

	appShortener, err := shortener.NewPolynomialHashShortener(10)
	if err != nil {
		log.Fatal(err)
	}

	linkInteractor := usecase.NewLinkInteractor(appShortener, appStorage)

	handler := http_server.NewHandler(appConfig.HttpServer, linkInteractor)
	server := http_server.NewServer(appConfig.HttpServer, handler)
	appCloser.Add(server.Shutdown)

	go server.Run()

	appCloser.Run()

}
