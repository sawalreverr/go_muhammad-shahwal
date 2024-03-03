package task

func SpinSlice(numbers []int) []int {
	n := len(numbers)
	mid := n / 2
	result := make([]int, n)

	for i := 0; i < mid; i++ {
		result[i*2] = numbers[i]
		result[i*2+1] = numbers[n-i-1]
	}

	if n%2 != 0 {
		result[n-1] = numbers[mid]
	}

	return result
}
