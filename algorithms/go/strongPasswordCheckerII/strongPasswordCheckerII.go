package strongPasswordCheckerII

import (
	"strings"
)

const otherString = "!@#$%^&*()-+"

func strongPasswordCheckerII(password string) bool {
	if len(password) < 8 {
		return false
	}

	var isLower, isUpper, isDigit, isSpecial bool

	for i, val := range password {
		if i != len(password)-1 && password[i] == password[i+1] {
			return false
		}
		if val >= 'a' && val <= 'z' {
			isLower = true
		} else if val >= 'A' && val <= 'Z' {
			isUpper = true
		} else if val >= '0' && val <= '9' {
			isDigit = true
		} else if strings.Contains(otherString, string(val)) {
			isSpecial = true
		}

	}
	return isLower && isUpper && isDigit && isSpecial
}
