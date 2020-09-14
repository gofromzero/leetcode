package program

import "testing"

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *TreeNode
		q *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "isSameTree",
			args: args{
				p: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 1,
					},
				},
				q: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 1,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			want: false,
		},
		{
			name: "isSameTree",
			args: args{
				p: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 1,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
				q: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 1,
					},
				},
			},
			want: false,
		},
		{
			name: "isSameTree",
			args: args{
				p: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 1,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
				q: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 1,
						Right: &TreeNode{
							Val: 2,
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
