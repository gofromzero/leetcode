package program

import "testing"

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lengthOfLIS case 1",
			args: args{
				nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			},
			want: 4,
		},
		{
			name: "lengthOfLIS case 2",
			args: args{
				nums: []int{2, 2},
			},
			want: 1,
		},
		{
			name: "lengthOfLIS case 3",
			args: args{
				nums: []int{1, 3, 6, 7, 9, 4, 10, 5, 6},
			},
			want: 6,
		},
		{
			name: "lengthOfLIS case 4",
			args: args{
				nums: []int{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLIS2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lengthOfLIS case 1",
			args: args{
				nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			},
			want: 4,
		},
		{
			name: "lengthOfLIS case 2",
			args: args{
				nums: []int{2, 2},
			},
			want: 1,
		},
		{
			name: "lengthOfLIS case 3",
			args: args{
				nums: []int{1, 3, 6, 7, 9, 4, 10, 5, 6},
			},
			want: 6,
		},
		{
			name: "lengthOfLIS case 4",
			args: args{
				nums: []int{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS2(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS2() = %v, want %v", got, tt.want)
			}
		})
	}
}
