package usecase

import "go-shortener/internal/entity"

type Storage interface {
	AddLink(link entity.Link) error
	GetLink(mapping string) (entity.Link, error)
}
