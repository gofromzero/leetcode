// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/6/9

package reverselist

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "reverseList",
			args: args{
				head: GenList(1, 2, 3, 4, 5),
			},
			want: GenList(5, 4, 3, 2, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
