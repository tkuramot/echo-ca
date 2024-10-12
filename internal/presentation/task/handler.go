package task

import (
	"github.com/labstack/echo/v4"
	taskApp "github/tkuramot/echo-practice/internal/application/task"
	echoRepo "github/tkuramot/echo-practice/internal/infrastructure/echo/repository"
	"github/tkuramot/echo-practice/internal/presentation/settings"
)

type Handler struct {
	findAllTasksUseCase *taskApp.FindAllTasksUseCase
	saveTaskUseCase     *taskApp.SaveTaskUseCase
}

func NewHandler(
	findAllTasksUseCase *taskApp.FindAllTasksUseCase,
	saveTaskUseCase *taskApp.SaveTaskUseCase,
) *Handler {
	return &Handler{
		findAllTasksUseCase: findAllTasksUseCase,
		saveTaskUseCase:     saveTaskUseCase,
	}
}

func (h *Handler) FindAllTasks(c echo.Context) error {
	ctx := c.Request().Context()
	sessionRepo := echoRepo.NewSessionRepository(c)
	userID, err := sessionRepo.UserID()
	if err != nil {
		return settings.ReturnStatusInternalServerError(c, err)
	}
	ts, err := h.findAllTasksUseCase.Run(ctx, userID)
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
		return settings.ReturnStatusInternalServerError(c, err)
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
