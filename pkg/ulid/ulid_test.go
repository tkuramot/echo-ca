package ulid

import "testing"

func TestIsValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid",
			args: args{
				s: NewULID(),
			},
			want: true,
		},
		{
			name: "invalid: empty",
			args: args{
				s: "",
			},
			want: false,
		},
		{
			name: "invalid: not enough length",
			args: args{
				s: "0123456789",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.args.s); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
