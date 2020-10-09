package program

import "testing"

func Test_findSecondMinimumValue(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "findSecondMinimumValue",
			args: args{
				root: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 5,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
			},
			want: 5,
		},
		{
			name: "findSecondMinimumValue",
			args: args{
				root: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 2,
					},
					Left: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 5,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
			},
			want: 5,
		},
		{
			name: "findSecondMinimumValue",
			args: args{
				root: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 2,
					},
					Left: &TreeNode{
						Val: 5,
					},
				},
			},
			want: 5,
		},
		{
			name: "findSecondMinimumValue",
			args: args{
				root: nil,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSecondMinimumValue(tt.args.root); got != tt.want {
				t.Errorf("findSecondMinimumValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
