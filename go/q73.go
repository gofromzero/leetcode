package program

func setZeroes(matrix [][]int) {
	mx := make([]int, 0)
	my := make([]int, 0)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				mx = append(mx, i)
				my = append(my, j)
			}
		}
	}
	for _, v := range mx {
		for i := 0; i < len(matrix[v]); i++ {
			matrix[v][i] = 0
		}
	}
	for _, v := range my {
		for i := 0; i < len(matrix); i++ {
			matrix[i][v] = 0
		}
	}

}
