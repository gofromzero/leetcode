package program

import "testing"

func Test_integerBreak(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "integerBreak",
			args: args{
				n: 3,
			},
			want: 2,
		},
		{
			name: "integerBreak",
			args: args{
				n: 4,
			},
			want: 4,
		},
		{
			name: "integerBreak",
			args: args{
				n: 5,
			},
			want: 6,
		},
		{
			name: "integerBreak",
			args: args{
				n: 6,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := integerBreak(tt.args.n); got != tt.want {
				t.Errorf("integerBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
