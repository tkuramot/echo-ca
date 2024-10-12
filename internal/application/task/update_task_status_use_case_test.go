package task

import (
	"errors"
	taskDomain "github/tkuramot/echo-practice/internal/domain/task"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUpdateTaskStatusUseCase_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := taskDomain.NewMockRepository(ctrl)
	uc := NewUpdateTaskStatusUseCase(mockRepo)

	tests := []struct {
		name     string
		taskID   string
		dto      UpdateTaskStatusUseCaseInputDto
		mockFunc func()
		wantErr  error
	}{
		{
			name:   "正常系",
			taskID: "taskID",
			dto:    UpdateTaskStatusUseCaseInputDto{Status: taskDomain.InProgress},
			mockFunc: func() {
				mockRepo.EXPECT().UpdateStatus(gomock.Any(), "taskID", taskDomain.InProgress).Return(nil)
			},
			wantErr: nil,
		},
		{
			name:     "異常系: 無効なステータス",
			taskID:   "taskID",
			dto:      UpdateTaskStatusUseCaseInputDto{Status: "invalid"},
			mockFunc: func() {},
			wantErr:  taskDomain.ErrInvalidStatus,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			err := uc.Run(nil, tt.taskID, tt.dto)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("want: %v, got: %v", tt.wantErr, err)
			}
		})
	}
}
