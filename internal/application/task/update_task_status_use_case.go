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
	Status taskDomain.Status
}

func (uc *UpdateTaskStatusUseCase) Run(
	ctx context.Context,
	taskID string,
	dto UpdateTaskStatusUseCaseInputDto,
) error {
	if taskDomain.IsValidStatus(dto.Status) == false {
		return taskDomain.ErrInvalidStatus
	}
	if err := uc.taskRepo.UpdateStatus(ctx, taskID, dto.Status); err != nil {
		return err
	}
	return nil
}
