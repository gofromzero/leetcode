package program

import (
	"reflect"
	"testing"
)

func Test_findMode(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		//{
		//	name: "findMode",
		//	args: args{
		//		root: &TreeNode{
		//			Val:   1,
		//			Right:  &TreeNode{
		//				Val:   2,
		//				Left:   &TreeNode{
		//					Val:   2,
		//				},
		//			},
		//		},
		//	},
		//	want: []int{2},
		//},
		{
			name: "findMode",
			args: args{
				root: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMode(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMode() = %v, want %v", got, tt.want)
			}
		})
	}
}
