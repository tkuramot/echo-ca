package session

import "testing"

func TestNewSession(t *testing.T) {
	type args struct {
		userID          string
		rememberMe      bool
		isAuthenticated bool
	}
	tests := []struct {
		name string
		args args
		want *Session
	}{
		{
			name: "success",
			args: args{
				userID:          "test",
				rememberMe:      true,
				isAuthenticated: true,
			},
			want: &Session{
				userID:          "test",
				rememberMe:      true,
				isAuthenticated: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSession(tt.args.userID, tt.args.rememberMe, tt.args.isAuthenticated)
			if got.userID != tt.want.userID {
				t.Errorf("NewSession() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSession_IsAuthenticated(t *testing.T) {
	type args struct {
		isAuthenticated bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				isAuthenticated: true,
			},
			want: true,
		},
		{
			name: "failed",
			args: args{
				isAuthenticated: false,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Session{
				isAuthenticated: tt.args.isAuthenticated,
			}
			if got := s.IsAuthenticated(); got != tt.want {
				t.Errorf("Session.IsAuthenticated() = %v, want %v", got, tt.want)
			}
		})
	}
}
