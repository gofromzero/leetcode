// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/6/8

package removenthfromend

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				head: GenList([]int{1, 2, 3, 4, 5}),
				n:    2,
			},
			want: GenList([]int{1, 2, 3, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
