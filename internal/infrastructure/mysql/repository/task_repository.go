package repository

import (
	"context"
	"database/sql"
	"github/tkuramot/echo-practice/internal/domain/task"
	taskDomain "github/tkuramot/echo-practice/internal/domain/task"
	"github/tkuramot/echo-practice/internal/infrastructure/mysql/db"
	"github/tkuramot/echo-practice/internal/infrastructure/mysql/db/dbgen"
)

type taskRepository struct{}

func NewTaskRepository() task.Repository {
	return &taskRepository{}
}

func (r *taskRepository) FindAll(ctx context.Context, filter taskDomain.Filter) ([]*task.Task, error) {
	query := db.GetQuery(ctx)
	ts, err := query.TaskFindAll(ctx, dbgen.TaskFindAllParams{
		UserID: sql.NullString{
			String: filter.UserID,
			Valid:  filter.UserID != "",
		},
		TaskStatus: dbgen.NullTasksStatus{
			TasksStatus: dbgen.TasksStatus(filter.Status),
			Valid:       filter.Status != "",
		},
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
	return err
}

func (r *taskRepository) Update(ctx context.Context, userID string, t *task.Task) error {
	query := db.GetQuery(ctx)
	err := query.TaskUpdate(ctx, dbgen.TaskUpdateParams{
		UserID:      userID,
		ID:          t.ID(),
		Title:       t.Title(),
		Description: t.Description(),
		Status:      dbgen.TasksStatus(t.Status()),
	})
	return err
}

func (r *taskRepository) UpdateStatus(ctx context.Context, userID, taskID string, status task.Status) error {
	query := db.GetQuery(ctx)
	err := query.TaskUpdateStatus(ctx, dbgen.TaskUpdateStatusParams{
		UserID: userID,
		ID:     taskID,
		Status: dbgen.TasksStatus(status),
	})
	return err
}
