package program

func plusOne(digits []int) []int {
	inc := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += inc
		if digits[i] < 10 {
			inc = 0
			break
		}
		inc = digits[i] / 10
		digits[i] %= 10
	}

	if inc != 0 {
		digits = append([]int{inc}, digits...)
	}

	return digits
}
