package memory_storage

import (
	"fmt"
	"go-shortener/internal/adapters/storage"
	"go-shortener/internal/entity"
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

	link, ok := ms.data[mapping]
	if !ok {
		return entity.Link{}, fmt.Errorf("MemoryStorage.GetLink(): %w", storage.ErrLinkNotFound)
	}

	return link, nil
}
