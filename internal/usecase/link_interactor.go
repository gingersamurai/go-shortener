package usecase

import (
	"context"
	"fmt"
	"go-shortener/internal/entity"
)

type Storage interface {
	AddLink(ctx context.Context, link entity.Link) error
	GetLink(ctx context.Context, mapping string) (entity.Link, error)
}

type Shortener interface {
	Shorten(source string) (string, error)
}

type LinkInteractor struct {
	shortener Shortener
	storage   Storage
}

func NewLinkInteractor(shortener Shortener, storage Storage) *LinkInteractor {
	return &LinkInteractor{
		shortener: shortener,
		storage:   storage,
	}
}

func (li *LinkInteractor) AddLink(source string) (string, error) {
	mapping, err := li.shortener.Shorten(source)
	if err != nil {
		return "", fmt.Errorf("LinkInteractor.AddLink(): %w", err)
	}

	link := entity.Link{
		Source:  source,
		Mapping: mapping,
	}
	err = li.storage.AddLink(context.TODO(), link)
	if err != nil {
		return "", fmt.Errorf("LinkInteractor.AddLink(): %w", err)
	}

	return mapping, nil
}

func (li *LinkInteractor) GetLink(mapping string) (string, error) {
	link, err := li.storage.GetLink(context.TODO(), mapping)
	if err != nil {
		return "", fmt.Errorf("LinkInteractor.GetLink(): %w", err)
	}

	return link.Source, nil
}
