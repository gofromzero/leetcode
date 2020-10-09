package program

import "testing"

func Test_deepestLeavesSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "deepestLeavesSum",
			args: args{},
			want: 0,
		},
		{
			name: "deepestLeavesSum",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 7,
							},
						},
						Right: &TreeNode{
							Val: 5,
						},
					},
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 6,
							Right: &TreeNode{
								Val: 8,
							},
						},
					},
				},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deepestLeavesSum(tt.args.root); got != tt.want {
				t.Errorf("deepestLeavesSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deepestLeavesSum2(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "deepestLeavesSum2",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 7,
							},
						},
						Right: &TreeNode{
							Val: 5,
						},
					},
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 6,
							Right: &TreeNode{
								Val: 8,
							},
						},
					},
				},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deepestLeavesSum2(tt.args.root); got != tt.want {
				t.Errorf("deepestLeavesSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
