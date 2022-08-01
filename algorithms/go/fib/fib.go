package fib

func fib(n int) int {
	l, r := 0, 1
	for i := 0; i < n; i++ {
		l, r = r, l+r
	}

	return l
}

//func fib(n int) int {
//	if n == 0 {
//		return 0
//	}
//	if n == 1 {
//		return 1
//	}
//	return fib(n-1) + fib(n-2)
//}
