package task

func GroupPrime(numbers []int) []any {
	var group []int
	var primes []int
	var returnGroup []any

	for _, v := range numbers {
		if Prime(v) {
			primes = append(primes, v)
		} else {
			group = append(group, v)
		}
	}

	if len(primes) != 0 {
		returnGroup = append(returnGroup, primes)
	}

	for _, v := range group {
		returnGroup = append(returnGroup, v)
	}

	return returnGroup
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
