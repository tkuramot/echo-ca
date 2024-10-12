package task

import taskDomain "github/tkuramot/echo-practice/internal/domain/task"

type updateTaskParams struct {
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Status      taskDomain.Status `json:"status"`
}

type updateTaskStatusParams struct {
	Status string `json:"status"`
}

type saveTaskParams struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
