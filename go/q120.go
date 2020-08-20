package program

import "math"

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}

	for i := 1; i < len(triangle); i++ {
		l := len(triangle[i])
		for j := 0; j < l; j++ {
			if j == 0 {
				triangle[i][j] = triangle[i-1][j] + triangle[i][j]
				continue
			}
			if j == l-1 {
				triangle[i][j] = triangle[i-1][j-1] + triangle[i][j]
				continue
			}
			triangle[i][j] = min(triangle[i-1][j], triangle[i-1][j-1]) + triangle[i][j]
		}
	}

	var result = math.MaxInt32
	for _, v := range triangle[len(triangle)-1] {
		result = min(result, v)
	}
	return result
}
