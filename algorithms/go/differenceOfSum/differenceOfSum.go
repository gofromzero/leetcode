package differenceOfSum

func differenceOfSum(nums []int) int {
	var sum int
	for _, val := range nums {
		sum += DifSum(val)
	}

	if sum < 0 {
		return -sum
	}
	return sum
}

func DifSum(num int) int {

	var sum int
	val := num
	for val != 0 {
		sum += val % 10
		val /= 10
	}
	return num - sum
}
