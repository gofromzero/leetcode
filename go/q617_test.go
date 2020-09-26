package program

import (
	"reflect"
	"testing"
)

func Test_mergeTrees(t *testing.T) {
	type args struct {
		t1 *TreeNode
		t2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "mergeTrees",
			args: args{
				t1: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 5}},
					Right: &TreeNode{Val: 2},
				},
				t2: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1, Right: &TreeNode{Val: 4}},
					Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 7}},
				},
			},
			want: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 4}},
				Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 7}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTrees(tt.args.t1, tt.args.t2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
