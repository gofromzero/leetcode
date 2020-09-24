package program

import "testing"

func Test_restoreString(t *testing.T) {
	type args struct {
		s       string
		indices []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "restoreString",
			args: args{
				s:       "codeleet",
				indices: []int{4, 5, 6, 7, 0, 2, 1, 3},
			},
			want: "leetcode",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreString(tt.args.s, tt.args.indices); got != tt.want {
				t.Errorf("restoreString() = %v, want %v", got, tt.want)
			}
		})
	}
}
