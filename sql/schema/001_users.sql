-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT UNIQUE NOT NULL
);

-- An up migration with queries to execute

-- +goose Down
DROP TABLE users;

-- A down migration to drop table