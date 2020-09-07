package program

import (
	"reflect"
	"testing"
)

func Test_numMovesStones(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "numMovesStones",
			args: args{
				a: 1,
				b: 2,
				c: 3,
			},
			want: []int{0, 0},
		},
		{
			name: "numMovesStones",
			args: args{
				a: 1,
				b: 2,
				c: 4,
			},
			want: []int{1, 1},
		},
		{
			name: "numMovesStones",
			args: args{
				a: 1,
				b: 3,
				c: 5,
			},
			want: []int{1, 2},
		},
		{
			name: "numMovesStones",
			args: args{
				a: 1,
				b: 4,
				c: 8,
			},
			want: []int{2, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numMovesStones(tt.args.a, tt.args.b, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numMovesStones() = %v, want %v", got, tt.want)
			}
		})
	}
}
