package program

import "testing"

func Test_findMin2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "findMin2",
			args: args{
				nums: []int{1, 3, 5},
			},
			want: 1,
		},
		{
			name: "findMin2",
			args: args{
				nums: []int{2, 2, 2, 0, 1},
			},
			want: 0,
		},
		{
			name: "findMin2",
			args: args{
				nums: []int{3, 3, 0, 3, 3, 3, 3, 3},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin2(tt.args.nums); got != tt.want {
				t.Errorf("findMin2() = %v, want %v", got, tt.want)
			}
		})
	}
}
