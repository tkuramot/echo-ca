-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD UNIQUE (email);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP INDEX email;
-- +goose StatementEnd
