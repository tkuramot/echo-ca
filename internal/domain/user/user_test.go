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
		password string
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
				password: "P4ssw0rd!",
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
				password: "P4ssw0rd!",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid email",
			args: args{
				email:    "test",
				nickname: "test",
				password: "P4ssw0rd!",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "password has no lower case",
			args: args{
				email:    "test@example.com",
				nickname: "test",
				password: "P4SSW0RD!",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "password has no upper case",
			args: args{
				email:    "test@example.com",
				nickname: "test",
				password: "p4ssw0rd!",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "password has no digit",
			args: args{
				email:    "test@example.com",
				nickname: "test",
				password: "Password!",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "password has no special character",
			args: args{
				email:    "test@example.com",
				nickname: "test",
				password: "P4ssw0rd",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "password is too short",
			args: args{
				email:    "test@example.com",
				nickname: "test",
				password: "P4ssw0!",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "password has multi-byte characters",
			args: args{
				email:    "test@example.com",
				nickname: "test",
				password: "P4ssw0rd!わ",
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(tt.args.email, tt.args.nickname, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			diff := cmp.Diff(
				got, tt.want,
				cmp.AllowUnexported(User{}),
				cmpopts.IgnoreFields(User{}, "id", "passwordDigest"),
			)
			if diff != "" {
				t.Errorf("NewUser() = %v, want %v. error is %s", got, tt.want, err)
			}
		})
	}
}
