package memory_storage

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-shortener/internal/entity"
	"go-shortener/internal/storage"
	"testing"
)

func TestMemoryStorage(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		ms := NewMemoryStorage()
		needLink := entity.Link{
			Source:  "bibaboba.com",
			Mapping: "abacaba",
		}
		ctx := context.Background()
		err := ms.AddLink(ctx, needLink)
		assert.NoError(t, err)

		gotLink, err := ms.GetLink(ctx, needLink.Mapping)
		assert.NoError(t, err)
		assert.Equal(t, gotLink, needLink)

	})

	t.Run("wrong mapping", func(t *testing.T) {
		ctx := context.Background()
		ms := NewMemoryStorage()
		needLink := entity.Link{
			Source:  "bibaboba.com",
			Mapping: "abacaba",
		}

		err := ms.AddLink(ctx, needLink)
		assert.NoError(t, err)

		_, err = ms.GetLink(ctx, needLink.Mapping+"fake suffix")
		assert.ErrorIs(t, err, storage.ErrLinkNotFound)

	})

	t.Run("link already exists", func(t *testing.T) {
		ctx := context.Background()
		ms := NewMemoryStorage()
		needLink := entity.Link{
			Source:  "bibaboba.com",
			Mapping: "abacaba",
		}

		err := ms.AddLink(ctx, needLink)
		assert.NoError(t, err)

		err = ms.AddLink(ctx, entity.Link{Source: needLink.Source})
		assert.ErrorIs(t, err, storage.ErrLinkAlreadyExists)
	})

	t.Run("with cancel", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		ms := NewMemoryStorage()
		needLink := entity.Link{
			Source:  "bibaboba.com",
			Mapping: "abacaba",
		}

		cancel()
		err := ms.AddLink(ctx, needLink)
		assert.Error(t, err)

	})
}
