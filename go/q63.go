package program

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}

	arr := make([]int, len(obstacleGrid[0]))
	for i, v := range obstacleGrid[0] {
		if v == 1 {
			break
		}
		arr[i] = 1
	}
	for i := 1; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 1 {
				arr[j] = 0
			} else if j != 0 {
				arr[j] += arr[j-1]
			}
		}
	}
	return arr[len(arr)-1]
}
