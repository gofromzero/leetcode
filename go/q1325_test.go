package program

import (
	"reflect"
	"testing"
)

func Test_removeLeafNodes(t *testing.T) {
	type args struct {
		root   *TreeNode
		target int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "removeLeafNodes",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 2,
						},
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 2,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
				},
				target: 2,
			},
			want: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeLeafNodes(tt.args.root, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeLeafNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
