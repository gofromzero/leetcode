package program

import (
	"testing"
)

func Test_searchBST(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "searchBST",
			args: args{
				root: nil,
				val:  0,
			},
			want: nil,
		},
		{
			name: "searchBST",
			args: args{
				root: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 3,
						},
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
				val: 2,
			},
			want: &TreeNode{
				Val: 2,
			},
		},
		{
			name: "searchBST",
			args: args{
				root: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 3,
						},
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
				val: 7,
			},
			want: &TreeNode{
				Val: 7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchBST(tt.args.root, tt.args.val)
			if got == nil {
				if got != tt.want {
					t.Errorf("searchBST() = %v, want %v", got, tt.want)
				}
			} else {
				if got.Val != tt.want.Val {
					t.Errorf("searchBST() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
