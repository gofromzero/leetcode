package program

import "testing"

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "searchMatrix",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 50},
				},
				target: 1,
			},
			want: true,
		},
		{
			name: "searchMatrix",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 50},
				},
				target: -1,
			},
			want: false,
		},
		{
			name: "searchMatrix",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 50},
				},
				target: 13,
			},
			want: false,
		},
		{
			name: "searchMatrix",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 50},
				},
				target: 333,
			},
			want: false,
		},
		{
			name: "searchMatrix",
			args: args{
				matrix: [][]int{
					{},
				},
				target: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchMatrixBetter(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "searchMatrix",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 50},
				},
				target: 1,
			},
			want: true,
		},
		{
			name: "searchMatrix",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 50},
				},
				target: -1,
			},
			want: false,
		},
		{
			name: "searchMatrix",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 50},
				},
				target: 13,
			},
			want: false,
		},
		{
			name: "searchMatrix",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 50},
				},
				target: 333,
			},
			want: false,
		},
		{
			name: "searchMatrix",
			args: args{
				matrix: [][]int{
					{},
				},
				target: 1,
			},
			want: false,
		},
		{
			name: "searchMatrix",
			args: args{
				matrix: [][]int{},
				target: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrixBetter(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrixBetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
