package preorder

import (
	"reflect"
	"testing"
)

func Test_preorder(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "",
			args: args{
				root: &Node{
					Val: 1,
					Children: []*Node{
						{
							Val: 3,
							Children: []*Node{
								{
									Val: 5,
								},
								{
									Val: 6,
								},
							},
						},
						{
							Val: 2,
						},
						{
							Val: 4,
						},
					},
				},
			},
			wantAns: []int{1, 3, 5, 6, 2, 4},
		},
		{
			name: "",
			args: args{
				root: nil,
			},
			wantAns: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := preorder(tt.args.root); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("preorder() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
