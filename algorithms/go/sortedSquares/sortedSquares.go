// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/6/1

package sortedSquares

import "sort"

func sortedSquares(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] *= nums[i]
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums
}

func sortedSquares2(nums []int) []int {
	result := make([]int, len(nums))
	num := len(nums) - 1
	l, r := 0, len(nums)-1

	for num >= 0 {
		if nums[l]*nums[l] > nums[r]*nums[r] {
			result[num] = nums[l] * nums[l]
			l++
		} else {
			result[num] = nums[r] * nums[r]
			r--
		}
		num--
	}

	return result
}
