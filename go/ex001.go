package program

import (
	"fmt"
	"math"
	"strings"
)

// 题目：大数打印

// 输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。

func printNumbers(n int) []int {
	len := int(math.Pow(10, float64(n)))
	res := make([]int, len-1)
	for i := 1; i < len; i++ {
		res[i-1] = i
	}
	return res
}

func printNumbers2(n int) {
	number := strings.Repeat("0", n)
	temp := []rune(number)
	for !incrementNumber(temp) {
		saveNumber(temp)
	}
}

func incrementNumber(number []rune) bool {
	var isBreak bool
	var carryFlag int32
	l := len(number)
	for i := l - 1; i >= 0; i-- {
		nSum := number[i] - '0' + carryFlag
		if i == l-1 {
			nSum++
		}
		if nSum >= 10 {
			if i == 0 {
				isBreak = true
			} else {
				nSum -= 10
				carryFlag = 1
				number[i] = '0' + nSum
			}
		} else {
			number[i] = nSum + '0'
			break
		}
	}

	return isBreak
}

func saveNumber(number []rune) {
	var isBegin0 = true
	for _, v := range number {
		if isBegin0 && v != '0' {
			isBegin0 = false
		}
		if !isBegin0 {
			fmt.Print(string(v))
		}
	}
	fmt.Println()
}
