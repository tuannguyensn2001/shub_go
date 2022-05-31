-- +goose Up
-- +goose StatementBegin
CREATE TABLE schedules (
    id INT NOT NULL AUTO_INCREMENT,
    title VARCHAR(255),
    start_time VARCHAR(255),
    day TIMESTAMP,
    link VARCHAR(255),
    class_id INT,
    type VARCHAR(255),
    status TINYINT DEFAULT 1,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS schedules;
-- +goose StatementEnd
