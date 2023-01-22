package isSumEqual

import "testing"

func Test_isSumEqual(t *testing.T) {
	type args struct {
		firstWord  string
		secondWord string
		targetWord string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				firstWord:  "acb",
				secondWord: "cba",
				targetWord: "cdb",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSumEqual(tt.args.firstWord, tt.args.secondWord, tt.args.targetWord); got != tt.want {
				t.Errorf("isSumEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
