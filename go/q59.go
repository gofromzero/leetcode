package program

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	lx, rx := 0, n-1
	ly, ry := 0, n-1
	var x = 1
	for {
		for i := lx; i <= rx; i++ {
			result[ly][i] = x
			x++
		}
		ly++
		if ly > ry {
			break
		}
		for i := ly; i <= ry; i++ {
			result[i][rx] = x
			x++
		}
		rx--
		for i := rx; i >= lx; i-- {
			result[ry][i] = x
			x++
		}
		ry--
		for i := ry; i >= ly; i-- {
			result[i][lx] = x
			x++
		}
		lx++
		if lx > rx {
			break
		}
	}
	return result
}
func generateMatrixBetter(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	for s, e, m := 0, n-1, 0; s <= e; s, e = s+1, e-1 {
		if s == e {
			m++
			result[s][e] = m
		}
		for j := s; j <= e-1; j++ {
			m++
			result[s][j] = m
		}
		for i := s; i <= e-1; i++ {
			m++
			result[i][e] = m
		}
		for j := e; j >= s+1; j-- {
			m++
			result[e][j] = m
		}
		for i := e; i >= s+1; i-- {
			m++
			result[i][s] = m
		}
	}
	return result
}
