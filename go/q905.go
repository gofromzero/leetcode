package program

func sortArrayByParity(A []int) []int {
	l := len(A)
	left, right := 0, l-1
	for left < right {
		if A[left]&1 != 0 {
			A[left], A[right] = A[right], A[left]
			right--
		} else {
			left++
		}
	}

	return A
}
