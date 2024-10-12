package task

import (
	"context"
	taskDomain "github/tkuramot/echo-practice/internal/domain/task"
)

type FindAllTasksUseCase struct {
	taskRepo taskDomain.Repository
}

func NewFindAllTasksUseCase(
	taskRepo taskDomain.Repository,
) *FindAllTasksUseCase {
	return &FindAllTasksUseCase{
		taskRepo: taskRepo,
	}
}

type FindAllTasksUseCaseInputDto struct {
	UserID string
	Status taskDomain.Status
}

type FindAllTasksUseCaseOutputDto struct {
	ID          string
	Title       string
	Description string
	Status      taskDomain.Status
}

func (uc *FindAllTasksUseCase) Run(ctx context.Context, dto FindAllTasksUseCaseInputDto) ([]*FindAllTasksUseCaseOutputDto, error) {
	ts, err := uc.taskRepo.FindAll(ctx, taskDomain.Filter{
		UserID: dto.UserID,
		Status: dto.Status,
	})
	if err != nil {
		return nil, err
	}

	var tasks []*FindAllTasksUseCaseOutputDto
	for _, t := range ts {
		tasks = append(tasks, &FindAllTasksUseCaseOutputDto{
			ID:          t.ID(),
			Title:       t.Title(),
			Description: t.Description(),
			Status:      t.Status(),
		})
	}
	return tasks, nil
}
