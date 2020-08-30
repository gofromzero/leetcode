package program

func singleNumber(nums []int) int {
	var result int
	for _, v := range nums {
		result ^= v
	}
	return result
}
