package user

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewUser(t *testing.T) {
	type args struct {
		email    string
		nickname string
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				email:    "test@example.com",
				nickname: "test",
			},
			want: &User{
				email:    "test@example.com",
				nickname: "test",
			},
			wantErr: false,
		},
		{
			name: "empty nickname",
			args: args{
				email:    "test@example.com",
				nickname: "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid email",
			args: args{
				email:    "test",
				nickname: "test",
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(tt.args.email, tt.args.nickname)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			diff := cmp.Diff(
				got, tt.want,
				cmp.AllowUnexported(User{}),
				cmpopts.IgnoreFields(User{}, "id"),
			)
			if diff != "" {
				t.Errorf("NewUser() = %v, want %v. error is %s", got, tt.want, err)
			}
		})
	}
}
