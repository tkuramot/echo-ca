package user

import (
	"context"
	"github.com/google/go-cmp/cmp"
	sessionDomain "github/tkuramot/echo-practice/internal/domain/session"
	userDomain "github/tkuramot/echo-practice/internal/domain/user"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestGetCurrentUserUseCase_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSessionRepo := sessionDomain.NewMockRepository(ctrl)
	mockUserRepo := userDomain.NewMockRepository(ctrl)
	uc := NewGetCurrentUseCase(mockUserRepo)

	tests := []struct {
		name            string
		sessionMockFunc func()
		userMockFunc    func()
		want            *GetCurrentUseCaseOutputDto
		wantErr         bool
	}{
		{
			name: "get current user by session and return dto",
			sessionMockFunc: func() {
				mockSessionRepo.
					EXPECT().
					UserID().
					Return("0123456789", nil)
			},
			userMockFunc: func() {
				mockUserRepo.
					EXPECT().
					FindByID(gomock.Any(), "0123456789").
					DoAndReturn(func(ctx context.Context, id string) (*userDomain.User, error) {
						return reconstructUser(
							"0123456789",
							"test@example.com",
							"test",
							"password_digest",
						)
					})
			},
			want: &GetCurrentUseCaseOutputDto{
				ID:       "0123456789",
				Email:    "test@example.com",
				Nickname: "test",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sessionMockFunc()
			tt.userMockFunc()
			got, err := uc.Run(context.Background(), mockSessionRepo)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCurrentUserUseCase.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Errorf("Run() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
