package memory_storage

import (
	"fmt"
	"go-shortener/internal/entity"
	"go-shortener/internal/storage"
	"sync"
)

type MemoryStorage struct {
	sync.RWMutex
	data map[string]entity.Link
}

func NewMemoryStorage() *MemoryStorage {
	data := make(map[string]entity.Link)

	return &MemoryStorage{data: data}
}

func (ms *MemoryStorage) AddLink(link entity.Link) error {
	ms.Lock()
	defer ms.Unlock()

	if _, ok := ms.data[link.Source]; ok {
		return fmt.Errorf("MemoryStorage.AddLink(): %w", storage.ErrLinkAlreadyExists)
	}

	ms.data[link.Source] = link
	return nil
}

func (ms *MemoryStorage) GetLink(mapping string) (entity.Link, error) {
	ms.RLock()
	defer ms.RUnlock()

	for _, link := range ms.data {
		if link.Mapping == mapping {
			return link, nil
		}
	}

	return entity.Link{}, fmt.Errorf("MemoryStorage.GetLink(): %w", storage.ErrLinkNotFound)
}