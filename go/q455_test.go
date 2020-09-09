package program

import "testing"

func Test_findContentChildren(t *testing.T) {
	type args struct {
		g []int
		s []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "findContentChildren",
			args: args{
				g: []int{10, 9, 8, 7},
				s: []int{5, 6, 7, 8},
			},
			want: 2,
		},
		{
			name: "findContentChildren",
			args: args{
				g: nil,
				s: []int{5, 6, 7, 8},
			},
			want: 0,
		},
		{
			name: "findContentChildren",
			args: args{
				g: []int{1},
				s: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findContentChildren(tt.args.g, tt.args.s); got != tt.want {
				t.Errorf("findContentChildren() = %v, want %v", got, tt.want)
			}
		})
	}
}
