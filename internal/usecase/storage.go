package usecase

import "go-shortener/internal/entity"

type Storage interface {
	AddLink(source, mapping string) (string, error)
	GetLink(source string) entity.Link
}
