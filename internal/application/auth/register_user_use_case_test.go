package auth

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"go.uber.org/mock/gomock"
	"testing"

	userDomain "github/tkuramot/echo-practice/internal/domain/user"
)

func TestRegisterUserUseCase_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepo := userDomain.NewMockUserRepository(ctrl)
	uc := NewRegisterUserUseCase(mockUserRepo)

	tests := []struct {
		name     string
		dto      RegisterUserUseCaseDto
		mockFunc func()
		want     *userDomain.User
		wantErr  bool
	}{
		{
			name: "register user and return user",
			dto: RegisterUserUseCaseDto{
				Email:    "test@example.com",
				Nickname: "test",
				Password: "P4ssw0rd!",
			},
			mockFunc: func() {
				mockUserRepo.
					EXPECT().
					Save(gomock.Any(), gomock.Any()).
					Return(nil)
				mockUserRepo.
					EXPECT().
					FindByEmail(gomock.Any(), gomock.Any()).
					DoAndReturn(func(ctx context.Context, email string) (*userDomain.User, error) {
						return reconstructUser(
							"0123456789",
							"test@example.com",
							"test",
							"password_digest",
						), nil
					})
			},
			want: reconstructUser(
				"0123456789",
				"test@example.com",
				"test",
				"password_digest",
			),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			got, err := uc.Run(context.Background(), tt.dto)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterUserUseCase.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			diff := cmp.Diff(
				got,
				tt.want,
				cmp.AllowUnexported(userDomain.User{}),
			)
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
) *userDomain.User {
	user, _ := userDomain.Reconstruct(
		id,
		email,
		nickname,
		passwordDigest,
	)
	return user
}
