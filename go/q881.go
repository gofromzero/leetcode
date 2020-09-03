package program

import (
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	left, right := 0, len(people)-1
	var num int
	for left <= right {
		if people[left]+people[right] <= limit {
			left++
		}
		right--
		num++
	}

	return num
}
