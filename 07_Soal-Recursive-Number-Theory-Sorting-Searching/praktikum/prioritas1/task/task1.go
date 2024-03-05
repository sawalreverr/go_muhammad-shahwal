package task

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

func FibX(n int) int {
	if n <= 0 {
		return 0
	}

	return Fibonacci(n) + FibX(n-1)
}
