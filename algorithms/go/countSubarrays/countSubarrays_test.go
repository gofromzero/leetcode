// Package countSubarrays
// @Author Spume
// @Date 2025-04-27 09:38

package countSubarrays

import "testing"

func Test_countSubarrays(t *testing.T) {
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
				nums: []int{1, 2, 1, 4, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubarrays(tt.args.nums); got != tt.want {
				t.Errorf("countSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
