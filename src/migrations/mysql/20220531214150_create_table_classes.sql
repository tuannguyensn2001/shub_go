-- +goose Up
-- +goose StatementBegin
CREATE TABLE classes
(
    id                 INT NOT NULL AUTO_INCREMENT,
    name               VARCHAR(255),
    code               VARCHAR(255),
    approve_student    TINYINT DEFAULT 0,
    prevent_quit_class TINYINT DEFAULT 0,
    show_mark          TINYINT DEFAULT 0,
    disable_newsfeed   TINYINT DEFAULT 0,
    subject_id         INT,
    grade_id           INT,
    user_id            INT,
    private_code       VARCHAR(255),
    created_at         TIMESTAMP,
    updated_at         TIMESTAMP,
    deleted_at         TIMESTAMP,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS classes;
-- +goose StatementEnd
