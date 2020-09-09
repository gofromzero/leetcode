package program

func isValidSudoku(board [][]byte) bool {
	xbox := [9][9]bool{}
	ybox := [9][9]bool{}
	dbox := [9][9]bool{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			v := board[i][j]
			if v == '.' {
				continue
			}
			v -= '1'
			if xbox[i][v] || ybox[j][v] || dbox[i/3*3+j/3][v] {
				return false
			}
			xbox[i][v] = true
			ybox[j][v] = true
			dbox[i/3*3+j/3][v] = true
		}
	}
	return true
}
