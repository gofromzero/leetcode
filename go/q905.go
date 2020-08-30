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

//func sortArrayByParity(A []int) []int {
//	j := 0
//	for i := range A {
//		if A[i]%2 == 0 {
//			A[j], A[i] = A[i], A[j]
//			j++
//		}
//	}
//	return A
//}
