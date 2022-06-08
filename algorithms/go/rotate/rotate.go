// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/6/8

package rotate

func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums[:len(nums)-k])
	reverse(nums[len(nums)-k:])
	reverse(nums)
}

func reverse(nums []int) {
	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
