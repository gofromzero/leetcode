package offer

import "testing"

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "maxDepth",
			args: args{
				root: &TreeNode{
					Val: 0,
				},
			},
			want: 1,
		},
		{
			name: "maxDepth",
			args: args{
				root: &TreeNode{
					Val: 0,
					Right: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 0,
						},
					},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
