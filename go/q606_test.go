package program

import "testing"

func Test_tree2str(t *testing.T) {
	type args struct {
		t *TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "tree2str",
			args: args{
				t: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: "1(2()(4))(3)",
		},
		{
			name: "tree2str",
			args: args{
				t: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: "1(2(4))(3)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tree2str(tt.args.t); got != tt.want {
				t.Errorf("tree2str() = %v, want %v", got, tt.want)
			}
		})
	}
}
