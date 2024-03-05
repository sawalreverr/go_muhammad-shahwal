package task

func Catalan(n int) int {
	if n <= 1 {
		return n
	}

	return Factorial(2*n) / (Factorial(n+1) * Factorial(n))
}

func Factorial(n int) int {
	if n <= 1 {
		return n
	}

	return n * Factorial(n-1)
}
