package lowestCommonAncestor

import (
	"reflect"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "",
			args: args{
				root: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 0},
						Right: &TreeNode{
							Val:   4,
							Left:  &TreeNode{Val: 3},
							Right: &TreeNode{Val: 5},
						},
					},
					Right: &TreeNode{
						Val:   8,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 9},
					},
				},
				p: &TreeNode{Val: 2},
				q: &TreeNode{Val: 6},
			},
			want: &TreeNode{Val: 6},
		},
		{
			name: "",
			args: args{
				root: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 0},
						Right: &TreeNode{
							Val:   4,
							Left:  &TreeNode{Val: 3},
							Right: &TreeNode{Val: 5},
						},
					},
					Right: &TreeNode{
						Val:   8,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 9},
					},
				},
				p: &TreeNode{Val: 2},
				q: &TreeNode{Val: 7},
			},
			want: &TreeNode{Val: 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got.Val, tt.want.Val) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
