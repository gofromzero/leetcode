// Package countInterestingSubarrays
// @Author Spume
// @Date 2025-04-25 09:41

package countInterestingSubarrays

import "testing"

func Test_countInterestingSubarrays(t *testing.T) {
	type args struct {
		nums   []int
		modulo int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "",
			args: args{
				nums:   []int{3, 2, 4},
				modulo: 2,
				k:      1,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countInterestingSubarrays(tt.args.nums, tt.args.modulo, tt.args.k); got != tt.want {
				t.Errorf("countInterestingSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
