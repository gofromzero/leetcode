package program

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	var i int
	for ; i < len(matrix); i++ {
		if matrix[i][0] > target {
			break
		}
	}
	if i == 0 {
		return false
	}
	i--
	for j := 0; j < len(matrix[i]); j++ {
		if matrix[i][j] > target {
			return false
		} else if matrix[i][j] == target {
			return true
		}
	}

	return false
}

func searchMatrixBetter(matrix [][]int, target int) bool {
	xl := len(matrix)
	if xl == 0 {
		return false
	}
	yl := len(matrix[0])
	if yl == 0 {
		return false
	}
	x := 0
	y := yl - 1
	for {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] < target {
			x++
		} else if matrix[x][y] > target {
			y--
		}

		if y < 0 || x >= xl {
			break
		}
	}

	return false
}
