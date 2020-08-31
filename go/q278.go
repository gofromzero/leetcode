package program

func firstBadVersion(n int) int {
	var left, right = 1, n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

var badVersion = 4

func isBadVersion(n int) bool {
	return n >= badVersion
}
