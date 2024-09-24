package password

import (
	"testing"
)

func TestVerify(t *testing.T) {
	password := "validPassword123"
	hashed, err := Hash(password)
	if err != nil {
		t.Fatalf("Hash() error = %v", err)
	}

	tests := []struct {
		name     string
		password string
		hashed   string
		wantErr  bool
	}{
		{"CorrectPassword", password, hashed, false},
		{"IncorrectPassword", "wrongPassword", hashed, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Verify(tt.password, tt.hashed)
			if (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
