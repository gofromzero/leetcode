package program

import (
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
			name: "lowestCommonAncestor",
			args: args{
				root: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 0,
						},
						Right: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 3,
							},
							Right: &TreeNode{
								Val: 5,
							},
						},
					},
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 8,
						},
					},
				},
				p: &TreeNode{Val: 2},
				q: &TreeNode{Val: 4},
			},
			want: &TreeNode{Val: 2},
		},
		{
			name: "lowestCommonAncestor",
			args: args{
				root: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 0,
						},
						Right: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 3,
							},
							Right: &TreeNode{
								Val: 5,
							},
						},
					},
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 8,
						},
					},
				},
				p: &TreeNode{Val: 8},
				q: &TreeNode{Val: 9},
			},
			want: &TreeNode{Val: 8},
		},
		{
			name: "lowestCommonAncestor",
			args: args{},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q)
			if tt.want == nil && got != nil {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
			if got != nil && got.Val != tt.want.Val {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
