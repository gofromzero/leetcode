package program

import "testing"

func Test_isSubPath(t *testing.T) {
	type args struct {
		head *ListNode
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "isSubPath",
			args: args{
				head: &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 8}}},
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:  4,
						Left: nil,
						Right: &TreeNode{
							Val: 2,
							Left: &TreeNode{
								Val:   1,
								Left:  nil,
								Right: nil,
							},
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 2,
							Left: &TreeNode{
								Val:   6,
								Left:  nil,
								Right: nil,
							},
							Right: &TreeNode{
								Val: 8,
								Left: &TreeNode{
									Val:   1,
									Left:  nil,
									Right: nil,
								},
								Right: &TreeNode{
									Val:   3,
									Left:  nil,
									Right: nil,
								},
							},
						},
						Right: nil,
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubPath(tt.args.head, tt.args.root); got != tt.want {
				t.Errorf("isSubPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
