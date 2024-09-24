package user

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"go.uber.org/mock/gomock"

	userDomain "github/tkuramot/echo-practice/internal/domain/user"
)

func TestFindUserUseCase_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepo := userDomain.NewMockUserRepository(ctrl)
	uc := NewFindUserUseCase(mockUserRepo)

	tests := []struct {
		name     string
		id       string
		mockFunc func()
		want     *FindUseCaseDto
		wantErr  bool
	}{
		{
			name: "get user by id and return dto",
			id:   "0123456789",
			mockFunc: func() {
				mockUserRepo.
					EXPECT().
					FindByID(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, id string) (*userDomain.User, error) {
						return reconstructUser(
							id,
							"test@example.com",
							"test",
							"password_digest",
						)
					})
			},
			want: &FindUseCaseDto{
				ID:       "0123456789",
				Email:    "test@example.com",
				Nickname: "test",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			got, err := uc.Run(context.Background(), tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindUserUseCase.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			diff := cmp.Diff(got, tt.want)
			if diff != "" {
				t.Errorf("Run() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func reconstructUser(
	id string,
	email string,
	nickname string,
	passwordDigest string,
) (*userDomain.User, error) {
	user, err := userDomain.Reconstruct(
		id,
		email,
		nickname,
		passwordDigest,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
