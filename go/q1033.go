package program

import "sort"

func numMovesStones(a int, b int, c int) []int {
	arr := []int{a, b, c}
	sort.Ints(arr)
	x := arr[1] - arr[0] - 1
	y := arr[2] - arr[1] - 1
	max := x + y
	var min int
	if x != 0 || y != 0 {
		if x > 1 && y > 1 {
			min = 2
		} else {
			min = 1
		}
	}

	return []int{min, max}
}
