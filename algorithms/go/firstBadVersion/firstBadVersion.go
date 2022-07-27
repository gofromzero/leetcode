package firstBadVersion

func isBadVersion(version int) bool {
	if version >= 4 {
		return true
	}
	return false
}

func firstBadVersion(n int) int {
	l, r := 1, n
	for l < r {
		mid := l + (r-l)>>1
		if isBadVersion(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return r
}
