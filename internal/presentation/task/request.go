package task

import taskDomain "github/tkuramot/echo-practice/internal/domain/task"

type findAllTasksParams struct {
	UserID string            `query:"user_id"`
	Status taskDomain.Status `query:"status"`
}

type updateTaskParams struct {
	TaskID      string            `param:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Status      taskDomain.Status `json:"status"`
}

type updateTaskStatusParams struct {
	TaskID string `param:"id"`
	Status string `json:"status"`
}

type saveTaskParams struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
