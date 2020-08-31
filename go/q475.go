package program

import (
	"math"
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	var res int
	n := len(heaters)
	sort.Ints(heaters)
	for _, v := range houses {
		left, right := 0, n
		for left < right {
			mid := left + (right-left)/2
			if v > heaters[mid] {
				left = mid + 1
			} else {
				right = mid
			}
		}
		var dist1 = math.MaxInt64
		if right != 0 {
			dist1 = int(math.Abs(float64(v - heaters[right-1])))
		}
		var dist2 = math.MaxInt64
		if right != n {
			dist2 = int(math.Abs(float64(v - heaters[right])))
		}
		res = max(res, min(dist1, dist2))
	}
	return res
}
