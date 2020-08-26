package program

import "testing"

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "isValidBST",
			args: args{
				root: &TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 1},
				},
			},
			want: false,
		},
		{
			name: "isValidBST",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: -1,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			want: true,
		},
		{
			name: "isValidBST",
			args: args{
				root: nil,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
