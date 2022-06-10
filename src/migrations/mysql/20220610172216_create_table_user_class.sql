-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_class
(
    user_id    INT NOT NULL,
    class_id   INT NOT NULL,
    role       TINYINT DEFAULT 1,
    created_at timestamp,
    updated_at timestamp,
    PRIMARY KEY (user_id,class_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_class;
-- +goose StatementEnd
