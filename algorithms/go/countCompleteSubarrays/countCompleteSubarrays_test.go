// Package countCompleteSubarrays
// @Author Spume
// @Date 2025-04-24 10:02

package countCompleteSubarrays

import "testing"

func Test_countCompleteSubarrays1(t *testing.T) {
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
				nums: []int{5, 5, 5, 5},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCompleteSubarrays(tt.args.nums); got != tt.want {
				t.Errorf("countCompleteSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
