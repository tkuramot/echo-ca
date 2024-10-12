package task

import (
	"context"
	taskDomain "github/tkuramot/echo-practice/internal/domain/task"
)

type UpdateTaskStatusUseCase struct {
	taskRepo taskDomain.Repository
}

func NewUpdateTaskStatusUseCase(
	taskRepo taskDomain.Repository,
) *UpdateTaskStatusUseCase {
	return &UpdateTaskStatusUseCase{
		taskRepo: taskRepo,
	}
}

type UpdateTaskStatusUseCaseInputDto struct {
	ID     string
	Status taskDomain.Status
}

func (uc *UpdateTaskStatusUseCase) Run(
	ctx context.Context,
	userID string,
	dto UpdateTaskStatusUseCaseInputDto,
) error {
	if !taskDomain.IsValidStatus(dto.Status) {
		return taskDomain.ErrInvalidStatus
	}
	if err := uc.taskRepo.UpdateStatus(ctx, userID, dto.ID, dto.Status); err != nil {
		return err
	}
	return nil
}
