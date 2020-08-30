package program

// map
// sum

func singleNumber2(nums []int) int {
	var result int
	var temp int
	for i := 0; i < 64; i++ {
		temp = 0

		for _, v := range nums {
			temp += v >> i & 1
		}

		result |= temp % 3 << i
	}
	return result
}

func singleNumber2Better(nums []int) int {
	var a, b int
	for _, v := range nums {
		b = (b ^ v) & (^a)
		a = (a ^ v) & (^b)
	}
	return b
}
