package program

import "bytes"

func rotateString(A string, B string) bool {
	if A == B {
		return true
	}
	if len(A) != len(B) {
		return false
	}
	for i := 0; i < len(A); i++ {
		if A[i:]+A[:i] == B {
			return true
		}
	}

	return false
}

func rotateStringBetter(a string, b string) bool {
	if len(a) > 100 || len(b) > 100 {
		return false
	}

	if a == b {
		return true
	}

	as := []byte(a)
	bs := []byte(b)

	var cnt int
	for cnt < len(as) {
		cnt++
		for i := 0; i < len(as)-1; i++ {
			as[i], as[i+1] = as[i+1], as[i]
		}

		if bytes.Equal(as, bs) {
			return true
		}
	}

	return false
}
