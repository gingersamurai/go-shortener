-- +goose Up
-- +goose StatementBegin
CREATE TABLE links (
    id SERIAL PRIMARY KEY,
    source TEXT UNIQUE NOT NULL,
    mapping TEXT UNIQUE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE links
-- +goose StatementEnd
