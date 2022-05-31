-- +goose Up
-- +goose StatementBegin
CREATE TABLE profiles
(
    id         INT NOT NULL AUTO_INCREMENT,
    name       VARCHAR(255),
    avatar     VARCHAR(255),
    user_id    INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS profiles;
-- +goose StatementEnd
