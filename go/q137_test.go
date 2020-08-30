package program

import "testing"

type args struct {
	nums []int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{
		name: "singleNumber2",
		args: args{
			nums: []int{2, 2, 3, 2},
		},
		want: 3,
	},
	{
		name: "singleNumber2",
		args: args{
			nums: []int{0, 1, 0, 1, 0, 1, 99},
		},
		want: 99,
	},
}

func Test_singleNumber2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber2(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_singleNumber2Better(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber2Better(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber2Better() = %v, want %v", got, tt.want)
			}
		})
	}
}
