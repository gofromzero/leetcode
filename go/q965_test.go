package program

import "testing"

func Test_isUnivalTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "isUnivalTree",
			args: args{
				root: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 0,
						},
						Right: &TreeNode{
							Val: 0,
						},
					},
					Right: &TreeNode{
						Val: 0,
						Right: &TreeNode{
							Val: 0,
						},
					},
				},
			},
			want: true,
		},
		{
			name: "isUnivalTree",
			args: args{
				root: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 0,
						},
						Right: &TreeNode{
							Val: 0,
						},
					},
					Right: &TreeNode{
						Val: 0,
						Right: &TreeNode{
							Val: 0,
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUnivalTree(tt.args.root); got != tt.want {
				t.Errorf("isUnivalTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
