-- +goose Up
-- +goose StatementBegin
CREATE TABLE posts
(
    id         INT NOT NULL AUTO_INCREMENT,
    content    LONGTEXT,
    class_id   INT,
    user_id    INT,
    is_show    TINYINT DEFAULT 1,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS posts;
-- +goose StatementEnd
