package program

import (
	"reflect"
	"testing"
)

func Test_generateListNode(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "GenerateListNode case 1",
			args: args{
				list: []int{1, 2, 3, 4, 5},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		},
		{
			name: "GenerateListNode case 2",
			args: args{
				list: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateListNode(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkListEqual(t *testing.T) {
	type args struct {
		left  []int
		right []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "CheckListEqual case 1",
			args: args{
				left:  []int{1, 2, 3, 4, 5},
				right: []int{1, 2, 3, 4, 5},
			},
			want: true,
		},
		{
			name: "CheckListEqual case 2",
			args: args{
				left:  []int{1, 2, 3, 4, 5},
				right: []int{1, 2, 3, 5, 5},
			},
			want: false,
		},
		{
			name: "CheckListEqual case 3",
			args: args{
				left:  []int{1, 2, 3, 5},
				right: []int{1, 2, 3, 4, 5},
			},
			want: false,
		},
		{
			name: "CheckListEqual case 4",
			args: args{
				left:  []int{1, 2, 3, 5},
				right: []int{1, 2, 3},
			},
			want: false,
		},
		{
			name: "CheckListEqual case 5",
			args: args{
				left:  []int{1, 2, 3},
				right: []int{1, 2, 3, 4},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			left := GenerateListNode(tt.args.left)
			right := GenerateListNode(tt.args.right)
			if got := CheckListEqual(left, right); got != tt.want {
				t.Errorf("CheckListEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
