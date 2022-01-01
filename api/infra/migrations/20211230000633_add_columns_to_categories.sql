-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

ALTER TABLE categories ADD COLUMN created datetime,
  ADD COLUMN updated datetime,
  ADD UNIQUE INDEX title_on_categories (title);
ALTER TABLE article_categories ADD COLUMN created datetime,
  ADD COLUMN updated datetime;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

ALTER TABLE categories DROP COLUMN created datetime,
  DROP COLUMN updated datetime,
  DROP UNIQUE INDEX title_on_categories (title);
ALTER TABLE article_categories DROP COLUMN created datetime,
  DROP COLUMN updated datetime;
