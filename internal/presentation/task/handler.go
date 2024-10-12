package task

import (
	"github.com/labstack/echo/v4"
	taskApp "github/tkuramot/echo-practice/internal/application/task"
	taskDomain "github/tkuramot/echo-practice/internal/domain/task"
	echoRepo "github/tkuramot/echo-practice/internal/infrastructure/echo/repository"
	"github/tkuramot/echo-practice/internal/presentation/settings"
)

type Handler struct {
	findAllTasksUseCase     *taskApp.FindAllTasksUseCase
	updateTaskUseCase       *taskApp.UpdateTaskUseCase
	updateTaskStatusUseCase *taskApp.UpdateTaskStatusUseCase
	saveTaskUseCase         *taskApp.SaveTaskUseCase
}

func NewHandler(
	findAllTasksUseCase *taskApp.FindAllTasksUseCase,
	updateTaskUseCase *taskApp.UpdateTaskUseCase,
	updateTaskStatusUseCase *taskApp.UpdateTaskStatusUseCase,
	saveTaskUseCase *taskApp.SaveTaskUseCase,
) *Handler {
	return &Handler{
		findAllTasksUseCase:     findAllTasksUseCase,
		updateTaskUseCase:       updateTaskUseCase,
		updateTaskStatusUseCase: updateTaskStatusUseCase,
		saveTaskUseCase:         saveTaskUseCase,
	}
}

func (h *Handler) FindAllTasks(c echo.Context) error {
	ctx := c.Request().Context()
	var params findAllTasksParams
	if err := c.Bind(&params); err != nil {
		return settings.ReturnStatusBadRequest(c, err)
	}
	ts, err := h.findAllTasksUseCase.Run(ctx, taskApp.FindAllTasksUseCaseInputDto{
		UserID: params.UserID,
		Status: params.Status,
	})
	if err != nil {
		return err
	}

	var tasks []taskResponseModel
	for _, t := range ts {
		tasks = append(tasks, taskResponseModel{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Status:      t.Status,
		})
	}
	return settings.ReturnStatusOK(c, findAllTasksResponse{
		Tasks: tasks,
	})
}

func (h *Handler) UpdateTask(c echo.Context) error {
	ctx := c.Request().Context()
	var params updateTaskParams
	if err := c.Bind(&params); err != nil {
		return settings.ReturnStatusBadRequest(c, err)
	}

	dto := taskApp.UpdateTaskUseCaseInputDto{
		ID:          params.TaskID,
		Title:       params.Title,
		Description: params.Description,
		Status:      params.Status,
	}
	sessionRepo := echoRepo.NewSessionRepository(c)
	userID, err := sessionRepo.UserID()
	if err != nil {
		return err
	}
	task, err := h.updateTaskUseCase.Run(ctx, userID, dto)
	if err != nil {
		return err
	}

	return settings.ReturnStatusOK(c, updateTaskResponse{
		Task: taskResponseModel{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
		},
	})
}

func (h *Handler) UpdateTaskStatus(c echo.Context) error {
	ctx := c.Request().Context()
	var params updateTaskStatusParams
	if err := c.Bind(&params); err != nil {
		return settings.ReturnStatusBadRequest(c, err)
	}

	dto := taskApp.UpdateTaskStatusUseCaseInputDto{
		ID:     params.TaskID,
		Status: taskDomain.Status(params.Status),
	}
	sessionRepo := echoRepo.NewSessionRepository(c)
	userID, err := sessionRepo.UserID()
	if err != nil {
		return err
	}
	err = h.updateTaskStatusUseCase.Run(ctx, userID, dto)
	if err != nil {
		return err
	}

	return settings.ReturnStatusNoContent(c)
}

func (h *Handler) SaveTask(c echo.Context) error {
	ctx := c.Request().Context()
	var params saveTaskParams
	if err := c.Bind(&params); err != nil {
		return settings.ReturnStatusBadRequest(c, err)
	}

	dto := taskApp.SaveTaskUseCaseInputDto{
		Title:       params.Title,
		Description: params.Description,
	}
	sessionRepo := echoRepo.NewSessionRepository(c)
	userID, err := sessionRepo.UserID()
	if err != nil {
		return err
	}
	task, err := h.saveTaskUseCase.Run(ctx, userID, dto)
	if err != nil {
		return err
	}

	return settings.ReturnStatusCreated(c, saveTaskResponse{
		Task: taskResponseModel{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
		},
	})
}
