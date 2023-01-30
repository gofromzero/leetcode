package giveGem

import "testing"

func Test_giveGem(t *testing.T) {
	type args struct {
		gem        []int
		operations [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				gem:        []int{3, 1, 2},
				operations: [][]int{{0, 2}, {2, 1}, {2, 0}},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := giveGem(tt.args.gem, tt.args.operations); got != tt.want {
				t.Errorf("giveGem() = %v, want %v", got, tt.want)
			}
		})
	}
}
