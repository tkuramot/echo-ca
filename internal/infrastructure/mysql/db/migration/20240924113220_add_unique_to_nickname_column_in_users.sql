-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD UNIQUE (nickname);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP INDEX nickname;
-- +goose StatementEnd
