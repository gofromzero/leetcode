// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/6/1

package movezeroes

func moveZeroes(nums []int) {
	l, r := 0, 0
	for r < len(nums) {
		if nums[r] != 0 {
			nums[l] = nums[r]
			l++
		}
		r++
	}
	for l < r {
		nums[l] = 0
		l++
	}
}
