package task

import (
	"strings"
	"testing"
)

func TestNewTask(t *testing.T) {
	type args struct {
		title       string
		description string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				title:       "タイトル",
				description: "説明",
			},
			wantErr: false,
		},
		{
			name: "タイトルが空文字の場合",
			args: args{
				title:       "",
				description: "説明",
			},
			wantErr: true,
		},
		{
			name: "タイトルが256文字以上の場合",
			args: args{
				title:       strings.Repeat("あ", 256),
				description: "説明",
			},
			wantErr: true,
		},
		{
			name: "説明が1001文字以上の場合",
			args: args{
				title:       "タイトル",
				description: strings.Repeat("あ", 1001),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewTask(tt.args.title, tt.args.description)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
