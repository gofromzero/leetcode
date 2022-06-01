// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/5/26

package addtwonumbers

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "addTwoNumbers ",
			args: args{
				l1: GenList([]int{2, 4, 3}),
				l2: GenList([]int{5, 6, 4}),
			},
			want: GenList([]int{7, 0, 8}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
