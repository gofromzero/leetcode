package program

import (
	"reflect"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		root *TreeNode
		key  int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "deleteNode",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 2},
						Right: &TreeNode{Val: 4},
					},
					Right: &TreeNode{
						Val: 6,

						Right: &TreeNode{Val: 7},
					},
				},
				key: 3,
			},
			want: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:  4,
					Left: &TreeNode{Val: 2},
				},
				Right: &TreeNode{
					Val: 6,

					Right: &TreeNode{Val: 7},
				},
			},
		},
		{
			name: "deleteNode",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:  4,
						Left: &TreeNode{Val: 2},
					},
					Right: &TreeNode{
						Val:   6,
						Right: &TreeNode{Val: 7},
					},
				},
				key: 6,
			},
			want: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:  4,
					Left: &TreeNode{Val: 2},
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		},
		{
			name: "deleteNode",
			args: args{
				root: nil,
				key:  6,
			},
			want: nil,
		},
		{
			name: "deleteNode",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:  4,
						Left: &TreeNode{Val: 2},
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
				key: 4,
			},
			want: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 6,
				},
			},
		},
		{
			name: "deleteNode",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 7,
						Left: &TreeNode{
							Val: 6,
						},
						Right: &TreeNode{
							Val: 10,
							Left: &TreeNode{
								Val:  9,
								Left: &TreeNode{Val: 8},
							},
							Right: &TreeNode{
								Val: 11,
							},
						},
					},
				},
				key: 7,
			},
			want: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 10,
						Left: &TreeNode{
							Val: 9,
						},
						Right: &TreeNode{
							Val: 11,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteNode(tt.args.root, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
