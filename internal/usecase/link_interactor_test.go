package usecase

import (
	"github.com/stretchr/testify/assert"
	"go-shortener/internal/adapters/storage"
	"go-shortener/internal/adapters/storage/memory_storage"
	"go-shortener/internal/usecase/shortener"
	"testing"
)

func TestLinkInteractor(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		phs, err := shortener.NewPolynomialHashShortener(5)
		assert.NoError(t, err)
		ms := memory_storage.NewMemoryStorage()
		li := NewLinkInteractor(phs, ms)

		source := "ozon.com/fintech"
		mapping, err := li.AddLink(source)
		assert.NoError(t, err)

		gotSource, err := li.GetLink(mapping)
		assert.NoError(t, err)
		assert.Equal(t, source, gotSource)
	})

	t.Run("link already exist", func(t *testing.T) {
		phs, err := shortener.NewPolynomialHashShortener(5)
		assert.NoError(t, err)
		ms := memory_storage.NewMemoryStorage()
		li := NewLinkInteractor(phs, ms)

		source := "ozon.com/fintech"
		_, err = li.AddLink(source)
		assert.NoError(t, err)

		source = "ozon.com/fintech"
		_, err = li.AddLink(source)
		assert.ErrorIs(t, err, storage.ErrLinkAlreadyExists)
	})

	t.Run("wrong link", func(t *testing.T) {
		phs, err := shortener.NewPolynomialHashShortener(5)
		assert.NoError(t, err)
		ms := memory_storage.NewMemoryStorage()
		li := NewLinkInteractor(phs, ms)

		source := "ozon.com/fintech"
		_, err = li.AddLink(source)
		assert.NoError(t, err)

		_, err = li.GetLink("wrong mapping string")
		assert.ErrorIs(t, err, storage.ErrLinkNotFound)
	})
}
