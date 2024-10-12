package task

import (
	"context"
	taskDomain "github/tkuramot/echo-practice/internal/domain/task"
)

type UpdateTaskUseCase struct {
	taskRepo taskDomain.Repository
}

func NewUpdateTaskUseCase(
	taskRepo taskDomain.Repository,
) *UpdateTaskUseCase {
	return &UpdateTaskUseCase{
		taskRepo: taskRepo,
	}
}

type UpdateTaskUseCaseInputDto struct {
	Title       string
	Description string
	Status      taskDomain.Status
}

type UpdateTaskUseCaseOutputDto struct {
	ID          string
	Title       string
	Description string
	Status      taskDomain.Status
}

func (uc *UpdateTaskUseCase) Run(
	ctx context.Context,
	userID string,
	taskID string,
	dto UpdateTaskUseCaseInputDto,
) (*UpdateTaskUseCaseOutputDto, error) {
	t, err := taskDomain.Reconstruct(taskID, dto.Title, dto.Description, dto.Status)
	if err != nil {
		return nil, err
	}
	if err := uc.taskRepo.Update(ctx, userID, t); err != nil {
		return nil, err
	}
	return &UpdateTaskUseCaseOutputDto{
		ID:          t.ID(),
		Title:       t.Title(),
		Description: t.Description(),
		Status:      t.Status(),
	}, nil
}
