package program

import "math/bits"

func hammingWeight2(num uint32) int {
	count := 0
	for num > 0 {
		num &= num - 1
		count++
	}
	return count
}

func hammingWeight(num uint32) int {
	return bits.OnesCount32(num)
}
