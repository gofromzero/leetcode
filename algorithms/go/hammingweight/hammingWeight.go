// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/6/8

package hammingweight

import "math/bits"

func hammingWeight(num uint32) int {
	return bits.OnesCount32(num)
}
