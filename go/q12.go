package program

import "strings"

var m = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

func intToRoman(num int) string {
	var str string
	if num > 1000 {
		n := num / 1000
		str += strings.Repeat(m[1000], n)
		num %= 1000
	}

	if num > 100 {
		n := num / 100
		switch {
		case n == 9:
			str += "CM"
			n = 0
		case n >= 5:
			str += "D"
			n -= 5
		case n == 4:
			str += "CD"
			n = 0
		}
		str += strings.Repeat(m[100], n)

		num %= 100
	}
	if num > 10 {
		n := num / 10
		switch {
		case n == 9:
			str += "XC"
			n = 0
		case n >= 5:
			str += "L"
			n -= 5
		case n == 4:
			str += "XL"
			n = 0
		}
		str += strings.Repeat(m[10], n)

		num %= 10
	}
	if num > 0 {
		switch {
		case num == 9:
			str += "IX"
			num = 0
		case num >= 5:
			str += "V"
			num -= 5
		case num == 4:
			str += "IV"
			num = 0
		}
		str += strings.Repeat(m[1], num)
	}

	return str
}
