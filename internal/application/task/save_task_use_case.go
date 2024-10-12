package task

import (
	"context"
	taskDomain "github/tkuramot/echo-practice/internal/domain/task"
)

type SaveTaskUseCase struct {
	taskRepo taskDomain.Repository
}

func NewSaveTaskUseCase(
	taskRepo taskDomain.Repository,
) *SaveTaskUseCase {
	return &SaveTaskUseCase{
		taskRepo: taskRepo,
	}
}

type SaveTaskUseCaseInputDto struct {
	Title       string
	Description string
}

type SaveTaskUseCaseOutputDto struct {
	ID          string
	Title       string
	Description string
	Status      taskDomain.Status
}

func (uc *SaveTaskUseCase) Run(ctx context.Context, userID string, dto SaveTaskUseCaseInputDto) (*SaveTaskUseCaseOutputDto, error) {
	t, err := taskDomain.NewTask(dto.Title, dto.Description)
	if err != nil {
		return nil, err
	}
	if err := uc.taskRepo.Save(ctx, userID, t); err != nil {
		return nil, err
	}
	return &SaveTaskUseCaseOutputDto{
		ID:          t.ID(),
		Title:       t.Title(),
		Description: t.Description(),
		Status:      t.Status(),
	}, nil
}
