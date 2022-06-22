package ishappy

import "fmt"

func isHappy(n int) bool {
	return checkHappy(n, n)
}

func checkHappy(s, n int) bool {
	if n == 1 {
		return true
	}
	if n < 6 {
		return false
	}
	var sum int
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	if sum == s {
		return false
	}
	fmt.Println(s, sum)
	return checkHappy(s, sum)
}

//func isHappy(n int) bool {
//
//	m := make(map[int]struct{})
//	m[n] = struct{}{}
//	for n != 1 {
//		n = sum(n)
//		fmt.Println("n :", n)
//		if _, ok := m[n]; ok {
//			return false
//		}
//		m[n] = struct{}{}
//	}
//	return true
//}
//
//func sum(n int) int {
//	var result int
//	for n != 0 {
//		temp := n % 10
//		result += temp * temp
//		n /= 10
//	}
//
//	return result
//}
