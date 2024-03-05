package task

func GetSequence(n int) int {
	if n <= 1 {
		return n
	}

	return n + GetSequence(n-1)
}
