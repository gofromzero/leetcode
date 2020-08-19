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
		want *listNode
	}{
		{
			name: "generateListNode case 1",
			args: args{
				list: []int{1, 2, 3, 4, 5},
			},
			want: &listNode{
				Val: 1,
				Next: &listNode{
					Val: 2,
					Next: &listNode{
						Val: 3,
						Next: &listNode{
							Val: 4,
							Next: &listNode{
								Val:  5,
								Next: &listNode{},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateListNode(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateListNode() = %v, want %v", got, tt.want)
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
			name: "checkListEqual case 1",
			args: args{
				left:  []int{1, 2, 3, 4, 5},
				right: []int{1, 2, 3, 4, 5},
			},
			want: true,
		},
		{
			name: "checkListEqual case 2",
			args: args{
				left:  []int{1, 2, 3, 4, 5},
				right: []int{1, 2, 3, 5, 5},
			},
			want: false,
		},
		{
			name: "checkListEqual case 3",
			args: args{
				left:  []int{1, 2, 3, 5},
				right: []int{1, 2, 3, 4, 5},
			},
			want: false,
		},
		{
			name: "checkListEqual case 4",
			args: args{
				left:  []int{1, 2, 3, 5},
				right: []int{1, 2, 3},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			left := generateListNode(tt.args.left)
			right := generateListNode(tt.args.right)
			if got := checkListEqual(left, right); got != tt.want {
				t.Errorf("checkListEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
