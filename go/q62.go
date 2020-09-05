package program

func uniquePaths(m int, n int) int {
	arr := make([]int, m)
	for i := 0; i < m; i++ {
		arr[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			arr[j] += arr[j-1]
		}
	}
	return arr[m-1]
}
