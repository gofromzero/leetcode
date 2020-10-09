package program

import "testing"

func Test_sumEvenGrandparent(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sumEvenGrandparent",
			args: args{
				root: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val: 7,
						Left: &TreeNode{
							Val: 2,
							Left: &TreeNode{
								Val: 9,
							},
						},
						Right: &TreeNode{
							Val: 7,
							Left: &TreeNode{
								Val: 1,
							},
							Right: &TreeNode{
								Val: 4,
							},
						},
					},
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 3,
							Right: &TreeNode{
								Val: 5,
							},
						},
					},
				},
			},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumEvenGrandparent(tt.args.root); got != tt.want {
				t.Errorf("sumEvenGrandparent() = %v, want %v", got, tt.want)
			}
		})
	}
}
