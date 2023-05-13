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

func (li *LinkInteractor) AddLink(ctx context.Context, source string) (string, error) {
	resultCh := make(chan struct {
		mapping string
		err     error
	})

	go func() {
		defer close(resultCh)
		mapping, err := li.shortener.Shorten(source)
		if err != nil {
			resultCh <- struct {
				mapping string
				err     error
			}{mapping: "", err: fmt.Errorf("LinkInteractor.AddLink(): %w", err)}
			return
		}

		link := entity.Link{
			Source:  source,
			Mapping: mapping,
		}
		err = li.storage.AddLink(ctx, link)
		if err != nil {
			resultCh <- struct {
				mapping string
				err     error
			}{mapping: "", err: fmt.Errorf("LinkInteractor.AddLink(): %w", err)}
			return
		}
		resultCh <- struct {
			mapping string
			err     error
		}{mapping: mapping, err: nil}
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case result := <-resultCh:
		return result.mapping, result.err
	}
}

func (li *LinkInteractor) GetLink(ctx context.Context, mapping string) (string, error) {
	resultCh := make(chan struct {
		source string
		err    error
	})

	go func() {
		defer close(resultCh)
		link, err := li.storage.GetLink(ctx, mapping)
		if err != nil {
			resultCh <- struct {
				source string
				err    error
			}{source: "", err: fmt.Errorf("LinkInteractor.GetLink(): %w", err)}
			return
		}
		resultCh <- struct {
			source string
			err    error
		}{source: link.Source, err: nil}
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case result := <-resultCh:
		return result.source, result.err
	}
}
