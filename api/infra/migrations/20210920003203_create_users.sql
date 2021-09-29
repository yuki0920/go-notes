-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE users (
    id int AUTO_INCREMENT,
    name varchar(20),
    password blob,
    created datetime,
    updated datetime,
    PRIMARY KEY(id)
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP TABLE users;
