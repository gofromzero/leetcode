package program

import (
	"fmt"
)

func getHint(secret string, guess string) string {
	l := len(secret)
	m := make([]int, 10)

	var A, B int
	for i := 0; i < l; i++ {
		if secret[i] == guess[i] {
			A++
		} else {
			m[secret[i]-'0']++
			m[guess[i]-'0']--
		}
	}

	B = l - A
	for _, v := range m {
		if v > 0 {
			B -= v
		}
	}

	return fmt.Sprintf("%dA%dB", A, B)
}
