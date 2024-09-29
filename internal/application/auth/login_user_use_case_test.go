package auth

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github/tkuramot/echo-practice/internal/domain/session"
	"go.uber.org/mock/gomock"
	"testing"

	errDomain "github/tkuramot/echo-practice/internal/domain/error"
	sessionDomain "github/tkuramot/echo-practice/internal/domain/session"
	"github/tkuramot/echo-practice/internal/domain/user"
	userDomain "github/tkuramot/echo-practice/internal/domain/user"
	pwd "github/tkuramot/echo-practice/pkg/password"
)

func TestLoginUserUseCase_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepo := user.NewMockRepository(ctrl)
	mockSessionRepo := session.NewMockRepository(ctrl)
	uc := NewLoginUserUseCase(mockUserRepo)

	tests := []struct {
		name            string
		dto             LoginUserUseCaseInputDto
		userMockFunc    func()
		sessionMockFunc func()
		want            *LoginUserUseCaseOutputDto
		wantErr         bool
	}{
		{
			name: "valid credentials",
			dto: LoginUserUseCaseInputDto{
				Email:      "test@example.com",
				Password:   "P4ssw0rd!",
				RememberMe: true,
			},
			userMockFunc: func() {
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
			sessionMockFunc: func() {
				mockSessionRepo.
					EXPECT().
					Save(reconstructSession("whatever", true, true)).
					Return(nil)
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
			userMockFunc: func() {
				mockUserRepo.
					EXPECT().
					FindByEmail(gomock.Any(), "wrong@example.com").
					Return(nil, errDomain.ErrNotFound)
			},
			sessionMockFunc: func() {
				mockSessionRepo.
					EXPECT().
					Save(gomock.Any()).
					Times(0)
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
			userMockFunc: func() {
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
			sessionMockFunc: func() {
				mockSessionRepo.
					EXPECT().
					Save(gomock.Any()).
					Times(0)
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.userMockFunc()
			tt.sessionMockFunc()
			got, err := uc.Run(context.Background(), mockSessionRepo, tt.dto)
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

func reconstructSession(userID string, isAuthenticated, rememberMe bool) *sessionDomain.Session {
	return sessionDomain.NewSession(userID, isAuthenticated, rememberMe)
}
