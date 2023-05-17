package usecase

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-shortener/internal/storage"
	"go-shortener/internal/storage/memory_storage"
	"go-shortener/pkg/polynomial_hash_shortener"
	"testing"
)

func TestLinkInteractor(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		ctx := context.Background()
		phs, err := polynomial_hash_shortener.NewPolynomialHashShortener(5)
		assert.NoError(t, err)
		ms := memory_storage.NewMemoryStorage()
		li := NewLinkInteractor(phs, ms)

		source := "ozon.com/fintech"
		mapping, err := li.AddLink(ctx, source)
		assert.NoError(t, err)

		gotSource, err := li.GetLink(ctx, mapping)
		assert.NoError(t, err)
		assert.Equal(t, source, gotSource)
	})

	t.Run("link already exist", func(t *testing.T) {
		ctx := context.Background()
		phs, err := polynomial_hash_shortener.NewPolynomialHashShortener(5)
		assert.NoError(t, err)
		ms := memory_storage.NewMemoryStorage()
		li := NewLinkInteractor(phs, ms)

		source := "ozon.com/fintech"
		_, err = li.AddLink(ctx, source)
		assert.NoError(t, err)

		source = "ozon.com/fintech"
		_, err = li.AddLink(ctx, source)
		assert.ErrorIs(t, err, storage.ErrLinkAlreadyExists)
	})

	t.Run("wrong link", func(t *testing.T) {
		ctx := context.Background()
		phs, err := polynomial_hash_shortener.NewPolynomialHashShortener(5)
		assert.NoError(t, err)
		ms := memory_storage.NewMemoryStorage()
		li := NewLinkInteractor(phs, ms)

		source := "ozon.com/fintech"
		_, err = li.AddLink(ctx, source)
		assert.NoError(t, err)

		_, err = li.GetLink(ctx, "wrong mapping string")
		assert.ErrorIs(t, err, storage.ErrLinkNotFound)
	})

	t.Run("with cancel", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		phs, err := polynomial_hash_shortener.NewPolynomialHashShortener(5)
		assert.NoError(t, err)
		ms := memory_storage.NewMemoryStorage()
		li := NewLinkInteractor(phs, ms)

		source := "ozon.com/fintech"
		cancel()
		_, err = li.AddLink(ctx, source)
		assert.Error(t, err)
	})
}
