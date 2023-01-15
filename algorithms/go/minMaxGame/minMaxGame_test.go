package minMaxGame

import "testing"

func Test_minMaxGame(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "minMaxGame",
			args: args{
				nums: []int{1, 3, 5, 2, 4, 8, 2, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMaxGame(tt.args.nums); got != tt.want {
				t.Errorf("minMaxGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
