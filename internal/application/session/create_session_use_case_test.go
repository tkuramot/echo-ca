package session

import (
	sessionDomain "github/tkuramot/echo-practice/internal/domain/session"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestCreateSessionUseCase_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	sessionRepo := sessionDomain.NewMockRepository(ctrl)

	type args struct {
		sessionRepo sessionDomain.Repository
		dto         CreateSessionUseCaseInputDto
	}
	tests := []struct {
		name     string
		args     args
		mockFunc func()
		wantErr  bool
	}{
		{
			name: "should save session",
			args: args{
				sessionRepo: sessionRepo,
				dto: CreateSessionUseCaseInputDto{
					IsAuthenticated: true,
					UserID:          "userID",
				},
			},
			mockFunc: func() {
				sessionRepo.
					EXPECT().
					Save(reconstructSession("userID", true)).
					Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := NewCreateSessionUseCase()
			tt.mockFunc()
			if err := uc.Run(tt.args.sessionRepo, tt.args.dto); (err != nil) != tt.wantErr {
				t.Errorf("CreateSessionUseCase.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func reconstructSession(userID string, isAuthenticated bool) *sessionDomain.Session {
	return sessionDomain.NewSession(userID, isAuthenticated)
}
