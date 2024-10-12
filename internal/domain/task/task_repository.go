//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE
package task

import "context"

type Repository interface {
	FindAll(ctx context.Context, userID string) ([]*Task, error)
	FindByID(ctx context.Context, taskID string) (*Task, error)
	FindByStatus(ctx context.Context, userID string, status Status) ([]*Task, error)
	Save(ctx context.Context, userID string, task *Task) error
	UpdateStatus(ctx context.Context, taskID string, status Status) error
}
