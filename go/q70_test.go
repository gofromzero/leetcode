package program

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "climbStairs case 1",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "climbStairs case 2",
			args: args{
				n: 2,
			},
			want: 2,
		},
		{
			name: "climbStairs case 3",
			args: args{
				n: 6,
			},
			want: 13,
		},
		{
			name: "climbStairs case 4",
			args: args{
				n: 10,
			},
			want: 89,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
