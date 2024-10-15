-- +goose Up
-- +goose StatementBegin
ALTER TABLE tasks
ADD COLUMN user_id VARCHAR(255) NOT NULL,
ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
