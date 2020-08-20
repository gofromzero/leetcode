package program

func climbStairs(n int) int {
	pre := 0
	cur := 1
	for i := 0; i < n; i++ {
		pre, cur = cur, pre+cur
	}

	return cur
}
