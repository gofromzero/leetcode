package program

import "testing"

func Test_isStraight(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "isStraight",
			args: args{
				nums: []int{2, 3, 4, 5, 6},
			},
			want: true,
		},
		{
			name: "isStraight",
			args: args{
				nums: []int{2, 2, 4, 5, 6},
			},
			want: false,
		},
		{
			name: "isStraight",
			args: args{
				nums: []int{0, 2, 4, 5, 6},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStraight(tt.args.nums); got != tt.want {
				t.Errorf("isStraight() = %v, want %v", got, tt.want)
			}
		})
	}
}
