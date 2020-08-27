package program

import "testing"

func Test_countNodes(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "countNodes",
			args: args{
				root: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 0,
						},
						Right: &TreeNode{
							Val: 0,
						},
					},
					Right: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 0,
						},
					},
				},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNodes(tt.args.root); got != tt.want {
				t.Errorf("countNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countNodes2(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "countNodes",
			args: args{
				root: nil,
			},
			want: 0,
		},
		{
			name: "countNodes",
			args: args{
				root: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 0,
						},
						Right: &TreeNode{
							Val: 0,
						},
					},
					Right: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 0,
						},
					},
				},
			},
			want: 6,
		},
		{
			name: "countNodes",
			args: args{
				root: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 0,
						},
						Right: &TreeNode{
							Val: 0,
						},
					},
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNodes2(tt.args.root); got != tt.want {
				t.Errorf("countNodes2() = %v, want %v", got, tt.want)
			}
		})
	}
}
