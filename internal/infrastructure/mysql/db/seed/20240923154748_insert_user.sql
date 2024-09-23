-- +goose Up
-- +goose StatementBegin
INSERT INTO users (
    id,
    email,
    nickname,
    created_at,
    updated_at
) VALUES (
    1,
    'test@example.com',
    'test',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users WHERE id = 1;
-- +goose StatementEnd
