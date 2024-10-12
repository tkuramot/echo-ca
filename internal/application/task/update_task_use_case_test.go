package task

import (
	"context"
	"github.com/google/go-cmp/cmp"
	taskDomain "github/tkuramot/echo-practice/internal/domain/task"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUpdateTaskUseCase_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := taskDomain.NewMockRepository(ctrl)
	uc := NewUpdateTaskUseCase(mockRepo)

	type args struct {
		ID          string
		Title       string
		Description string
		Status      taskDomain.Status
	}
	tests := []struct {
		name     string
		args     args
		mockFunc func()
		want     *UpdateTaskUseCaseOutputDto
		wantErr  bool
	}{
		{
			name: "success",
			args: args{
				ID:          "id",
				Title:       "title",
				Description: "description",
				Status:      taskDomain.InProgress,
			},
			mockFunc: func() {
				mockRepo.EXPECT().Update(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			want: &UpdateTaskUseCaseOutputDto{
				ID:          "taskID",
				Title:       "title",
				Description: "description",
				Status:      taskDomain.InProgress,
			},
			wantErr: false,
		},
		{
			name: "empty title",
			args: args{
				ID:          "taskID",
				Title:       "",
				Description: "description",
				Status:      taskDomain.InProgress,
			},
			mockFunc: func() {},
			want:     nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			got, err := uc.Run(context.Background(), "userID", "taskID", UpdateTaskUseCaseInputDto{
				Title:       tt.args.Title,
				Description: tt.args.Description,
				Status:      tt.args.Status,
			})
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateTaskUseCase.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			diff := cmp.Diff(tt.want, got)
			if diff != "" {
				t.Errorf("UpdateTaskUseCase.Run() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
