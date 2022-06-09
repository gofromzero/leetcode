// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/6/9

package mergetwolists

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "mergeTwoLists",
			args: args{
				list1: GenList(1, 2, 4),
				list2: GenList(1, 3, 4),
			},
			want: GenList(1, 1, 2, 3, 4, 4),
		},
		{
			name: "mergeTwoLists",
			args: args{
				list1: GenList(1, 2, 4),
				list2: GenList(1, 3),
			},
			want: GenList(1, 1, 2, 3, 4),
		},
		{
			name: "mergeTwoLists",
			args: args{
				list1: GenList(1, 2),
				list2: GenList(1, 3, 4),
			},
			want: GenList(1, 1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
