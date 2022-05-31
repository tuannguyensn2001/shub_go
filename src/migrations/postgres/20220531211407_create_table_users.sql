-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id         SERIAL NOT NULL,
    username   TEXT,
    email      TEXT,
    password   TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY id
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
