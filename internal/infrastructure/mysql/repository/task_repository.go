package repository

import (
	"context"
	"github/tkuramot/echo-practice/internal/domain/task"
	"github/tkuramot/echo-practice/internal/infrastructure/mysql/db"
	"github/tkuramot/echo-practice/internal/infrastructure/mysql/db/dbgen"
)

type taskRepository struct{}

func NewTaskRepository() task.Repository {
	return &taskRepository{}
}

func (r *taskRepository) FindAll(ctx context.Context, userID string) ([]*task.Task, error) {
	query := db.GetQuery(ctx)
	ts, err := query.TaskFindAll(ctx, userID)
	if err != nil {
		return nil, err
	}
	var tasks []*task.Task
	for _, t := range ts {
		td, err := task.Reconstruct(
			t.ID,
			t.Title,
			t.Description,
			task.Status(t.Status),
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, td)
	}
	return tasks, nil
}

func (r *taskRepository) FindByStatus(ctx context.Context, userID string, status task.Status) ([]*task.Task, error) {
	query := db.GetQuery(ctx)
	ts, err := query.TaskFindByStatus(ctx, dbgen.TaskFindByStatusParams{
		UserID: userID,
		Status: dbgen.TasksStatus(status),
	})
	if err != nil {
		return nil, err
	}
	var tasks []*task.Task
	for _, t := range ts {
		td, err := task.Reconstruct(
			t.ID,
			t.Title,
			t.Description,
			task.Status(t.Status),
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, td)
	}
	return tasks, nil
}

func (r *taskRepository) Save(ctx context.Context, userID string, t *task.Task) error {
	query := db.GetQuery(ctx)
	err := query.TaskInsert(ctx, dbgen.TaskInsertParams{
		ID:          t.ID(),
		Title:       t.Title(),
		Description: t.Description(),
		Status:      dbgen.TasksStatus(t.Status()),
	})
	if err != nil {
		return err
	}
	err = query.UserTaskInsert(ctx, dbgen.UserTaskInsertParams{
		UserID: userID,
		TaskID: t.ID(),
	})
	if err != nil {
		return err
	}
	return err
}

func (r *taskRepository) UpdateStatus(ctx context.Context, taskID string, status task.Status) error {
	query := db.GetQuery(ctx)
	err := query.TaskUpdateStatus(ctx, dbgen.TaskUpdateStatusParams{
		ID:     taskID,
		Status: dbgen.TasksStatus(status),
	})
	return err
}
