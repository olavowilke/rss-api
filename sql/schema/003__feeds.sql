-- +goose Up
CREATE TABLE feeds
(
    id         UUID PRIMARY KEY,
    created_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name       TEXT        NOT NULL,
    url        TEXT UNIQUE NOT NULL,
    user_id    UUID REFERENCES users (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;
