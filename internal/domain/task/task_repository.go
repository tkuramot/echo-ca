//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE
package task

import (
	"context"
)

type Filter struct {
	UserID string
	Status Status
}

type Repository interface {
	FindAll(ctx context.Context, filter Filter) ([]*Task, error)
	FindByStatus(ctx context.Context, userID string, status Status) ([]*Task, error)
	Save(ctx context.Context, userID string, task *Task) error
	Update(ctx context.Context, userID string, task *Task) error
	UpdateStatus(ctx context.Context, userID, taskID string, status Status) error
}
