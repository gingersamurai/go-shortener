package postgres_storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"go-shortener/internal/entity"
	"go-shortener/internal/storage"
)

type PostgresStorage struct {
	conn *pgx.Conn
}

func NewPostgresStorage(databaseURL string) (*PostgresStorage, error) {
	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		return nil, err
	}
	err = conn.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return &PostgresStorage{conn: conn}, nil
}

func (ps *PostgresStorage) AddLink(ctx context.Context, link entity.Link) error {
	requestSQL := "SELECT source FROM links WHERE source = $1"
	row := ps.conn.QueryRow(ctx, requestSQL, link.Source)

	var source string
	err := row.Scan(&source)
	if err == nil {
		return fmt.Errorf("PostgresStorage.AddLink(): %w", storage.ErrLinkAlreadyExists)
	}
	requestSQL = "INSERT INTO links(source, mapping) VALUES ($1, $2)"
	_, err = ps.conn.Exec(ctx, requestSQL, link.Source, link.Mapping)
	if err != nil {
		return err
	}
	return nil

}

func (ps *PostgresStorage) GetLink(ctx context.Context, mapping string) (entity.Link, error) {
	requestSQL := "SELECT source, mapping FROM links WHERE mapping = $1"
	row := ps.conn.QueryRow(ctx, requestSQL, mapping)
	link := entity.Link{}
	err := row.Scan(&link.Source, &link.Mapping)
	if err != nil {
		return entity.Link{}, fmt.Errorf("PostgresStorage.GetLink(): %w", storage.ErrLinkNotFound)
	}
	return link, nil
}
