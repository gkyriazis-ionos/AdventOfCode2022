package main

import "testing"

func Test_findMaxBlock(t *testing.T) {

	maxBlockAsFinal := `
1
2

1
2
3`

	tests := []struct {
		name    string
		content string
		want    int
		wantErr bool
	}{
		{
			name:    "Stefan found a bug",
			content: maxBlockAsFinal,
			want:    6,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findMaxBlock(tt.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("findMaxBlock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("findMaxBlock() want: %d, got: %d", tt.want, got)
			}
		})
	}
}

func Test_findSumOfMaxBlock(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findSumOfMaxBlock(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("findSumOfMaxBlock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("findSumOfMaxBlock() got = %v, want %v", got, tt.want)
			}
		})
	}
}
