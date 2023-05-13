package memory_storage

import (
	"context"
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

func (ms *MemoryStorage) AddLink(ctx context.Context, link entity.Link) error {
	resultCh := make(chan error)
	go func() {
		ms.Lock()
		defer ms.Unlock()
		if _, ok := ms.data[link.Source]; ok {
			resultCh <- fmt.Errorf("MemoryStorage.AddLink(): %w", storage.ErrLinkAlreadyExists)
			return
		}

		ms.data[link.Source] = link
		resultCh <- nil
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case result := <-resultCh:
		return result
	}
}

func (ms *MemoryStorage) GetLink(ctx context.Context, mapping string) (entity.Link, error) {
	resultCh := make(chan struct {
		link entity.Link
		err  error
	})

	go func() {
		ms.RLock()
		defer ms.RUnlock()

		for _, link := range ms.data {
			if link.Mapping == mapping {
				//return link, nil
				resultCh <- struct {
					link entity.Link
					err  error
				}{link: link, err: nil}
				return
			}

		}
		resultCh <- struct {
			link entity.Link
			err  error
		}{link: entity.Link{}, err: fmt.Errorf("MemoryStorage.GetLink(): %w", storage.ErrLinkNotFound)}
		return
	}()

	select {
	case <-ctx.Done():
		return entity.Link{}, ctx.Err()
	case result := <-resultCh:
		return result.link, result.err
	}
	//select {
	//case <-ctx.Done():
	//	return entity.Link{}, ctx.Err()
	//default:
	//	defer ms.RUnlock()
	//	ms.RLock()
	//
	//	for _, link := range ms.data {
	//		if link.Mapping == mapping {
	//			return link, nil
	//		}
	//	}
	//	return entity.Link{}, fmt.Errorf("MemoryStorage.GetLink(): %w", storage.ErrLinkNotFound)
	//}

}
