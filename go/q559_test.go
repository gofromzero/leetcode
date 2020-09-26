package program

import "testing"

func Test_maxNDepth(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "maxNDepth",
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
			want: 3,
		},
		{
			name: "maxNDepth",
			args: args{
				root: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxNDepth(tt.args.root); got != tt.want {
				t.Errorf("maxNDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
