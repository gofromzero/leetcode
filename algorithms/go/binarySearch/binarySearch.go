// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/5/27

package binarySearch

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	var middle int

	for left <= right {
		middle = left + (right-left)>>1
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return -1
}
