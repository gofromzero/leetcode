// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/5/27

package binarySearch

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "search",
			args: args{
				nums:   []int{1, 2, 3, 4, 5},
				target: 3,
			},
			want: 2,
		},
		{
			name: "search",
			args: args{
				nums:   []int{1, 2, 3, 4, 5},
				target: 8,
			},
			want: -1,
		},
		{
			name: "search",
			args: args{
				nums:   []int{1, 2, 3, 4, 5},
				target: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
