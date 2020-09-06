package program

import "testing"

func Test_uniquePathsWithObstacles(t *testing.T) {
	type args struct {
		obstacleGrid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "uniquePaths",
			args: args{
				obstacleGrid: [][]int{
					{0},
					{1},
				},
			},
			want: 0,
		},
		{
			name: "uniquePaths",
			args: args{
				obstacleGrid: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{0, 0, 0},
				},
			},
			want: 2,
		},
		{
			name: "uniquePaths",
			args: args{
				obstacleGrid: [][]int{
					{0, 0, 0, 0},
					{0, 1, 0, 0},
					{0, 0, 0, 0},
				},
			},
			want: 4,
		},
		{
			name: "uniquePaths",
			args: args{
				obstacleGrid: [][]int{
					{0, 1, 0, 0},
					{0, 1, 0, 0},
					{0, 0, 0, 0},
				},
			},
			want: 1,
		},

		{
			name: "uniquePaths",
			args: args{
				obstacleGrid: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsWithObstacles(tt.args.obstacleGrid); got != tt.want {
				t.Errorf("uniquePathsWithObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}
