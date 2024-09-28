package auth

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"go.uber.org/mock/gomock"
	"testing"

	errDomain "github/tkuramot/echo-practice/internal/domain/error"
	"github/tkuramot/echo-practice/internal/domain/user"
	userDomain "github/tkuramot/echo-practice/internal/domain/user"
	pwd "github/tkuramot/echo-practice/pkg/password"
)

func TestLoginUserUseCase_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepo := user.NewMockUserRepository(ctrl)
	uc := NewLoginUserUseCase(mockUserRepo)

	tests := []struct {
		name     string
		dto      LoginUserUseCaseInputDto
		mockFunc func()
		want     *LoginUserUseCaseOutputDto
		wantErr  bool
	}{
		{
			name: "valid credentials",
			dto: LoginUserUseCaseInputDto{
				Email:    "test@example.com",
				Password: "P4ssw0rd!",
			},
			mockFunc: func() {
				mockUserRepo.
					EXPECT().
					FindByEmail(gomock.Any(), "test@example.com").
					Return(
						reconstructUser(
							"whatever",
							"test@example.com",
							"test",
							"P4ssw0rd!",
						),
					)
			},
			want: &LoginUserUseCaseOutputDto{
				ID:       "whatever",
				Email:    "test@example.com",
				Nickname: "test",
			},
			wantErr: false,
		},
		{
			name: "wrong email",
			dto: LoginUserUseCaseInputDto{
				Email:    "wrong@example.com",
				Password: "P4ssw0rd!",
			},
			mockFunc: func() {
				mockUserRepo.
					EXPECT().
					FindByEmail(gomock.Any(), "wrong@example.com").
					Return(nil, errDomain.ErrNotFound)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "wrong password",
			dto: LoginUserUseCaseInputDto{
				Email:    "test@example.com",
				Password: "wrong",
			},
			mockFunc: func() {
				mockUserRepo.
					EXPECT().
					FindByEmail(gomock.Any(), "test@example.com").
					Return(
						reconstructUser(
							"whatever",
							"test@example.com",
							"test",
							"P4ssw0rd!",
						),
					)
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			got, err := uc.Run(context.Background(), tt.dto)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginUserUseCase.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			diff := cmp.Diff(
				got,
				tt.want,
				cmpopts.IgnoreFields(LoginUserUseCaseOutputDto{}, "ID"),
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
	password string,
) (*userDomain.User, error) {
	digest, err := pwd.Hash(password)
	u, err := userDomain.Reconstruct(
		id,
		email,
		nickname,
		digest,
	)
	if err != nil {
		return nil, err
	}
	return u, nil
}
