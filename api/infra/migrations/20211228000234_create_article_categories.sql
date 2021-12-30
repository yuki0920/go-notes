-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE article_categories (
    id int AUTO_INCREMENT,
    article_id int,
    category_id int,
    PRIMARY KEY(id),
    FOREIGN KEY(article_id) REFERENCES articles(id),
    FOREIGN KEY(category_id) REFERENCES categories(id)
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP TABLE article_categories;
