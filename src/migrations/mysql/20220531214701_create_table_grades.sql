-- +goose Up
-- +goose StatementBegin
CREATE TABLE grades
(
    id         INT NOT NULL AUTO_INCREMENT,
    name       VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS grades;
-- +goose StatementEnd
