package task

import taskDomain "github/tkuramot/echo-practice/internal/domain/task"

type findAllTasksResponse struct {
	Tasks []taskResponseModel `json:"tasks"`
}

type updateTaskResponse struct {
	Task taskResponseModel `json:"task"`
}

type saveTaskResponse struct {
	Task taskResponseModel `json:"task"`
}

type taskResponseModel struct {
	ID          string            `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Status      taskDomain.Status `json:"status"`
}
