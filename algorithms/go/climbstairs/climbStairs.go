// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/6/8

package climbstairs

func climbStairs(n int) int {
	l := 1
	r := 1
	for i := 1; i < n; i++ {
		l, r = l+r, l

	}
	return l
}
