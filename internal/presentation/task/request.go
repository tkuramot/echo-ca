package task

type updateTaskStatusParams struct {
	Status string `json:"status"`
}

type saveTaskParams struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
