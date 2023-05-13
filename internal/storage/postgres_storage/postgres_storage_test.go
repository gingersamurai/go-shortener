package postgres_storage

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-shortener/internal/entity"
	"go-shortener/internal/storage"
	"testing"
)

func TestPostgresStorage(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		ctx := context.Background()
		ps, err := NewPostgresStorage("host=localhost user=postgres password=12345678")
		assert.NoError(t, err)

		_, _ = ps.conn.Exec(context.Background(), "DELETE FROM links")

		link := entity.Link{
			Source:  "bibaboba.com",
			Mapping: "aaaa",
		}
		err = ps.AddLink(ctx, link)
		assert.NoError(t, err)

		gotSource, err := ps.GetLink(ctx, link.Mapping)
		assert.NoError(t, err)
		assert.Equal(t, gotSource, link)
	})
	t.Run("same link", func(t *testing.T) {
		ctx := context.Background()
		ps, err := NewPostgresStorage("host=localhost user=postgres password=12345678")
		assert.NoError(t, err)

		_, _ = ps.conn.Exec(context.Background(), "DELETE FROM links")

		link := entity.Link{
			Source:  "bibaboba.com",
			Mapping: "aaaa",
		}
		err = ps.AddLink(ctx, link)
		assert.NoError(t, err)

		err = ps.AddLink(ctx, link)
		assert.ErrorIs(t, err, storage.ErrLinkAlreadyExists)

	})

	t.Run("bad mapping", func(t *testing.T) {
		ctx := context.Background()
		ps, err := NewPostgresStorage("host=localhost user=postgres password=12345678")
		assert.NoError(t, err)

		_, _ = ps.conn.Exec(context.Background(), "DELETE FROM links")

		link := entity.Link{
			Source:  "bibaboba.com",
			Mapping: "aaaa",
		}
		err = ps.AddLink(ctx, link)
		assert.NoError(t, err)

		_, err = ps.GetLink(ctx, "bbbb")
		assert.ErrorIs(t, err, storage.ErrLinkNotFound)

	})

	t.Run("with cancel", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		ps, err := NewPostgresStorage("host=localhost user=postgres password=12345678")
		assert.NoError(t, err)

		_, _ = ps.conn.Exec(context.Background(), "DELETE FROM links")

		link := entity.Link{
			Source:  "bibaboba.com",
			Mapping: "aaaa",
		}
		cancel()
		err = ps.AddLink(ctx, link)
		assert.Error(t, err)

	})

}
