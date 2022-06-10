-- +goose Up
-- +goose StatementBegin
CREATE TABLE comments (
    id INT NOT NULL AUTO_INCREMENT,
    content LONGTEXT,
    post_id INT,
    user_id INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS comments;
-- +goose StatementEnd
