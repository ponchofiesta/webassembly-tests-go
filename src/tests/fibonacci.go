package tests

func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	a := 0
	b := 1
	for n--; n > 0; n-- {
		a += b
		a, b = b, a
	}
	return b
}
