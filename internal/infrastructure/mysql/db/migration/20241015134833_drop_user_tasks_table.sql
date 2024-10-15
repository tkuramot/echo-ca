-- +goose Up
-- +goose StatementBegin
DROP TABLE  user_tasks;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
CREATE TABLE user_tasks (
    user_id VARCHAR(255) NOT NULL,
    task_id VARCHAR(255) NOT NULL,
    PRIMARY KEY (user_id, task_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (task_id) REFERENCES tasks(id)
);
-- +goose StatementEnd
