package task

func PrimeSum(numbers []int) int {
	var total int

	for _, x := range numbers {
		if Prime(x) {
			total += x
		}
	}

	return total
}

func Prime(n int) bool {
	if n == 1 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	}

	for i := 3; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
