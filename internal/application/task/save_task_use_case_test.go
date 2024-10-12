package task

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	taskDomain "github/tkuramot/echo-practice/internal/domain/task"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestSaveTaskUseCase_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockTaskRepo := taskDomain.NewMockRepository(ctrl)
	uc := NewSaveTaskUseCase(mockTaskRepo)

	tests := []struct {
		name     string
		dto      SaveTaskUseCaseInputDto
		mockFunc func()
		want     *SaveTaskUseCaseOutputDto
		wantErr  bool
	}{
		{
			name: "save task and return task",
			dto: SaveTaskUseCaseInputDto{
				Title:       "test",
				Description: "test",
			},
			mockFunc: func() {
				mockTaskRepo.
					EXPECT().
					Save(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil)
			},
			want: &SaveTaskUseCaseOutputDto{
				ID:          "whatever",
				Title:       "test",
				Description: "test",
				Status:      taskDomain.NotStarted,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			got, err := uc.Run(context.Background(), "whatever", tt.dto)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveTaskUseCase.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			diff := cmp.Diff(
				got,
				tt.want,
				cmpopts.IgnoreFields(SaveTaskUseCaseOutputDto{}, "ID"),
			)
			if diff != "" {
				t.Errorf("SaveTaskUseCase.Run() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
