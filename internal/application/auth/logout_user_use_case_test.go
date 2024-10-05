package auth

import (
	sessionDomain "github/tkuramot/echo-practice/internal/domain/session"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestLogoutUserUseCase_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSessionRepo := sessionDomain.NewMockRepository(ctrl)
	uc := NewLogoutUserUseCase()

	tests := []struct {
		name     string
		mockFunc func()
		wantErr  bool
	}{
		{
			name: "success",
			mockFunc: func() {
				mockSessionRepo.EXPECT().Delete().Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			if err := uc.Run(mockSessionRepo); (err != nil) != tt.wantErr {
				t.Errorf("LogoutUserUseCase.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
