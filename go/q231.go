package program

func isPowerOfTwoBetter(n int) bool {
	return n > 0 && n&(n-1) == 0
}

func isPowerOfTwo(n int) bool {
	temp := 1
	for temp < n {
		temp = temp << 1
	}

	return temp == n
}
