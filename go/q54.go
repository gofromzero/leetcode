package program

var changes = [][]int{
	{0, 1},
	{1, 0},
	{-1, 0},
	{0, -1},
}

func spiralOrder(matrix [][]int) []int {
	h := len(matrix)
	if h == 0 {
		return nil
	}
	var result []int
	l, r := 0, len(matrix[0])-1 // 左右
	h, b := 0, len(matrix)-1    // 上下
	for {
		for i := l; i <= r; i++ {
			result = append(result, matrix[h][i])
		}
		h++
		if h > b {
			break
		}
		for i := h; i <= b; i++ {
			result = append(result, matrix[i][r])
		}
		r--
		if l > r {
			break
		}
		for i := r; i >= l; i-- {
			result = append(result, matrix[b][i])
		}
		b--
		if h > b {
			break
		}
		for i := b; i >= h; i-- {
			result = append(result, matrix[i][l])
		}
		l++
		if l > r {
			break
		}
	}
	return result
}
