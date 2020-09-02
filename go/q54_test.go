package program

import (
	"reflect"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "spiralOrder",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: []int{
				1, 2, 3, 6, 9, 8, 7, 4, 5,
			},
		},
		{
			name: "spiralOrder",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
			},
			want: []int{
				1, 2, 3, 6, 5, 4,
			},
		},
		{
			name: "spiralOrder",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
				},
			},
			want: []int{
				1, 2, 3,
			},
		},
		{
			name: "spiralOrder",
			args: args{
				matrix: [][]int{
					{1},
					{2},
				},
			},
			want: []int{
				1, 2,
			},
		},
		{
			name: "spiralOrder",
			args: args{
				matrix: [][]int{
					{1, 2},
					{3, 4},
					{5, 6},
				},
			},
			want: []int{
				1, 2, 4, 6, 5, 3,
			},
		},
		{
			name: "spiralOrder",
			args: args{
				matrix: [][]int{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralOrder(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
