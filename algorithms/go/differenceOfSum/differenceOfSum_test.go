package differenceOfSum

import "testing"

func Test_differenceOfSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 15, 6, 3},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := differenceOfSum(tt.args.nums); got != tt.want {
				t.Errorf("differenceOfSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
