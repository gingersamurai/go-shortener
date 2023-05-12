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
		ps, err := NewPostgresStorage("host=localhost user=postgres password=12345678")
		assert.NoError(t, err)

		_, _ = ps.conn.Exec(context.Background(), "DELETE FROM links")

		link := entity.Link{
			Source:  "bibaboba.com",
			Mapping: "aaaa",
		}
		err = ps.AddLink(link)
		assert.NoError(t, err)

		gotSource, err := ps.GetLink(link.Mapping)
		assert.NoError(t, err)
		assert.Equal(t, gotSource, link)
	})
	t.Run("same link", func(t *testing.T) {
		ps, err := NewPostgresStorage("host=localhost user=postgres password=12345678")
		assert.NoError(t, err)

		_, _ = ps.conn.Exec(context.Background(), "DELETE FROM links")

		link := entity.Link{
			Source:  "bibaboba.com",
			Mapping: "aaaa",
		}
		err = ps.AddLink(link)
		assert.NoError(t, err)

		err = ps.AddLink(link)
		assert.ErrorIs(t, err, storage.ErrLinkAlreadyExists)

	})

	t.Run("bad mapping", func(t *testing.T) {
		ps, err := NewPostgresStorage("host=localhost user=postgres password=12345678")
		assert.NoError(t, err)

		_, _ = ps.conn.Exec(context.Background(), "DELETE FROM links")

		link := entity.Link{
			Source:  "bibaboba.com",
			Mapping: "aaaa",
		}
		err = ps.AddLink(link)
		assert.NoError(t, err)

		_, err = ps.GetLink("bbbb")
		assert.ErrorIs(t, err, storage.ErrLinkNotFound)

	})
}
