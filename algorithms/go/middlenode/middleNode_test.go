// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/6/1

package middlenode

import (
	"testing"
)

func Test_middleNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "middleNode",
			args: args{
				head: GenList([]int{1, 2, 3, 4, 5}),
			},
			want: GenList([]int{3, 4, 5}),
		},
		{
			name: "middleNode",
			args: args{
				head: GenList([]int{1, 2, 3, 4, 5, 6}),
			},
			want: GenList([]int{4, 5, 6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); !EqualListNode(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
