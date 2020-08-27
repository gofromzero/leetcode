package program

import (
	"reflect"
	"testing"
)

func Test_pruneTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "pruneTree",
			args: args{
				root: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val: 0,
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			want: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 0,
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pruneTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pruneTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
