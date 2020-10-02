package program

import "testing"

func Test_isSubtree(t *testing.T) {
	type args struct {
		s *TreeNode
		t *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "isSubtree",
			args: args{
				s: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				t: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			want: true,
		},
		{
			name: "isSubtree",
			args: args{
				s: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				t: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: false,
		},
		{
			name: "isSubtree",
			args: args{
				s: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				t: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 1,
					},
				},
			},
			want: false,
		},
		{
			name: "isSubtree",
			args: args{
				s: nil,
				t: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 1,
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubtree(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
