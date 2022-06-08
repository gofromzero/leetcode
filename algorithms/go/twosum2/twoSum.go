// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/6/8

package twosum2

func twoSum(numbers []int, target int) []int {

	l := 0
	r := len(numbers) - 1
	for {
		ans := numbers[l] + numbers[r]
		if ans > target {
			r--
		} else if ans < target {
			l++
		} else {
			return []int{l + 1, r + 1}
		}
	}

	return []int{}
}
