package program

import (
	"reflect"
	"testing"
)

func Test_postorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{
			name: "postorderTraversal",
			args: args{
				root: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3,
						},
					},
				},
			},
			wantRes: []int{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := postorderTraversal(tt.args.root); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("postorderTraversal() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
