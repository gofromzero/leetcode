package program

func gameOfLife(board [][]int) {
	var dx = []int{-1, 0, 1, -1, 1, -1, 0, 1}
	var dy = []int{-1, -1, -1, 0, 0, 1, 1, 1}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			var sum int
			for k := 0; k < 8; k++ {
				nx := i + dx[k]
				ny := j + dy[k]
				if nx >= 0 && nx < len(board) && ny >= 0 && ny < len(board[i]) {
					sum += board[nx][ny] & 1
				}
			}
			if board[i][j] == 1 {
				if sum == 2 || sum == 3 {
					board[i][j] |= 2
				}
			} else {
				if sum == 3 {
					board[i][j] |= 2
				}
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			board[i][j] >>= 1
		}
	}
}
