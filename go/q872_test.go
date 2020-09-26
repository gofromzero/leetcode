package program

import "testing"

func Test_leafSimilar(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "leafSimilar",
			args: args{
				root1: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 6,
						},
						Right: &TreeNode{
							Val: 2,
							Left: &TreeNode{
								Val: 7,
							},
							Right: &TreeNode{
								Val: 4,
							},
						},
					},
					Right: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 9,
						},
						Right: &TreeNode{
							Val: 8,
						},
					},
				},
				root2: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 9,
							Left: &TreeNode{
								Val: 7,
							},
							Right: &TreeNode{
								Val: 4,
							},
						},
						Right: &TreeNode{
							Val: 8,
							Left: &TreeNode{
								Val: 9,
							},
							Right: &TreeNode{
								Val: 8,
							},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "leafSimilar",
			args: args{},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leafSimilar(tt.args.root1, tt.args.root2); got != tt.want {
				t.Errorf("leafSimilar() = %v, want %v", got, tt.want)
			}
		})
	}
}
