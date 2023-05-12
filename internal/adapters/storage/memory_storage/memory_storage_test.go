package memory_storage

import (
	"github.com/stretchr/testify/assert"
	"go-shortener/internal/adapters/storage"
	"go-shortener/internal/entity"
	"testing"
)

func TestMemoryStorage(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		ms := NewMemoryStorage()
		needLink := entity.Link{
			Source:  "bibaboba.com",
			Mapping: "abacaba",
		}

		err := ms.AddLink(needLink)
		assert.NoError(t, err)

		gotLink, err := ms.GetLink(needLink.Mapping)
		assert.NoError(t, err)
		assert.Equal(t, gotLink, needLink)

	})

	t.Run("wrong mapping", func(t *testing.T) {
		ms := NewMemoryStorage()
		needLink := entity.Link{
			Source:  "bibaboba.com",
			Mapping: "abacaba",
		}

		err := ms.AddLink(needLink)
		assert.NoError(t, err)

		_, err = ms.GetLink(needLink.Mapping + "fake suffix")
		assert.ErrorIs(t, err, storage.ErrLinkNotFound)

	})

	t.Run("link already exists", func(t *testing.T) {
		ms := NewMemoryStorage()
		needLink := entity.Link{
			Source:  "bibaboba.com",
			Mapping: "abacaba",
		}

		err := ms.AddLink(needLink)
		assert.NoError(t, err)

		err = ms.AddLink(entity.Link{Source: needLink.Source})
		assert.ErrorIs(t, err, storage.ErrLinkAlreadyExists)
	})
}
