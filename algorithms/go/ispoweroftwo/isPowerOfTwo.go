// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/6/8

package ispoweroftwo

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return n&(n-1) == 0
}
