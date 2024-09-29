package user

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"go.uber.org/mock/gomock"
	"testing"

	userDomain "github/tkuramot/echo-practice/internal/domain/user"
)

func TestRegisterUserUseCase_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepo := userDomain.NewMockRepository(ctrl)
	uc := NewRegisterUserUseCase(mockUserRepo)

	tests := []struct {
		name     string
		dto      RegisterUserUseCaseInputDto
		mockFunc func()
		want     *RegisterUserUseCaseOutputDto
		wantErr  bool
	}{
		{
			name: "register user and return user",
			dto: RegisterUserUseCaseInputDto{
				Email:    "test@example.com",
				Nickname: "test",
				Password: "P4ssw0rd!",
			},
			mockFunc: func() {
				mockUserRepo.
					EXPECT().
					Save(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			want: &RegisterUserUseCaseOutputDto{
				ID:       "whatever",
				Email:    "test@example.com",
				Nickname: "test",
			},
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
				cmpopts.IgnoreFields(RegisterUserUseCaseOutputDto{}, "ID"),
			)
			if diff != "" {
				t.Errorf("Run() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
