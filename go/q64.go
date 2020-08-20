package program

func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				grid[i][j] += grid[i][j-1]
				continue
			}
			if j == 0 {
				grid[i][j] += grid[i-1][j]
				continue
			}
			grid[i][j] += min(grid[i][j-1], grid[i-1][j])
		}
	}

	return grid[m-1][n-1]
}

func minPathSumBetter(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}

	dp := make([]int, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 || j == 0 {
				dp[j] += grid[j][i]
				if j > 0 {
					dp[j] += dp[j-1]
				}
				continue
			}
			dp[j] = min(dp[j-1], dp[j]) + grid[j][i]
		}
	}

	return dp[m-1]
}
