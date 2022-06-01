// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/5/27

package searchinsert

func searchInsert(nums []int, target int) int {
	i, j := 0, len(nums)-1
	res := len(nums)
	for i <= j {
		m := (j-i)/2 + i
		if nums[m] >= target {
			res = m
			j = m - 1
		} else {
			i = m + 1
		}
	}
	return res
}
