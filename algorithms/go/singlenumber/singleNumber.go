// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/6/9

package singlenumber

func singleNumber(nums []int) int {
	result := nums[0]
	for _, v := range nums[1:] {
		result ^= v
	}
	return result
}
