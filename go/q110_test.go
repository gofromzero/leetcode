package program

import "testing"

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *treeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "isBalanced case 1",
			args: args{
				root: &treeNode{
					Val: 3,
					Left: &treeNode{
						Val: 9,
					},
					Right: &treeNode{
						Val: 20,
						Left: &treeNode{
							Val: 15,
						},
						Right: &treeNode{
							Val: 7,
						},
					},
				},
			},
			want: true,
		},
		{
			name: "isBalanced case 2",
			args: args{
				root: &treeNode{
					Val: 1,
					Left: &treeNode{
						Val: 2,
						Left: &treeNode{
							Val: 3,
							Left: &treeNode{
								Val: 4,
							},
							Right: &treeNode{
								Val: 4,
							},
						},
						Right: &treeNode{
							Val: 3,
						},
					},
					Right: &treeNode{
						Val: 2,
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
