package task

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	taskDomain "github/tkuramot/echo-practice/internal/domain/task"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestFindAllTasksUseCase_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := taskDomain.NewMockRepository(ctrl)
	uc := NewFindAllTasksUseCase(mockRepo)

	tests := []struct {
		name     string
		mockFunc func()
		want     []*FindAllTasksUseCaseOutputDto
		wantErr  bool
	}{
		{
			name: "find all tasks",
			mockFunc: func() {
				mockRepo.
					EXPECT().
					FindAll(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, userID string) ([]*taskDomain.Task, error) {
						return []*taskDomain.Task{
							reconstructTask("test1", "test1"),
							reconstructTask("test2", "test2"),
						}, nil
					})
			},
			want: []*FindAllTasksUseCaseOutputDto{
				{
					Title:       "test1",
					Description: "test1",
					Status:      "not_started",
				},
				{
					Title:       "test2",
					Description: "test2",
					Status:      "not_started",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			got, err := uc.Run(context.Background(), "whatever")
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAllTasksUseCase.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			diff := cmp.Diff(
				got,
				tt.want,
				cmp.AllowUnexported(FindAllTasksUseCaseOutputDto{}, taskDomain.Task{}),
				cmpopts.IgnoreFields(FindAllTasksUseCaseOutputDto{}, "ID"),
			)
			if diff != "" {
				t.Errorf("Run() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func reconstructTask(title, description string) *taskDomain.Task {
	t, _ := taskDomain.NewTask(title, description)
	return t
}
