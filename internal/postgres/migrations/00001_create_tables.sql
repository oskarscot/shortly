-- +goose Up
CREATE TABLE IF NOT EXISTS links (
    id TEXT PRIMARY KEY NOT NULL,
    code TEXT NOT NULL,
    url TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS links;
