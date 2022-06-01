// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/5/25

package twosum

func twoSum(nums []int, target int) []int {
	var l = make(map[int]int)
	for i, v := range nums {
		if raw, ok := l[v]; ok {
			return []int{raw, i}
		}
		l[target-v] = i
	}
	return nil

}
