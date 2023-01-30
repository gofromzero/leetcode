package giveGem

import "math"

func giveGem(gem []int, operations [][]int) int {

	var half int
	for _, xy := range operations {
		x, y := xy[0], xy[1]
		half = gem[x] / 2
		gem[x] -= half
		gem[y] += half
	}
	low := math.MaxInt
	high := 0
	for _, val := range gem {
		if val > high {
			high = val
		}
		if val < low {
			low = val
		}
	}
	return high - low
}
