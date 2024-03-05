package task

func PrimeFactors(n int) []int {
	var faktor []int

	for n%2 == 0 {
		faktor = append(faktor, 2)
		n /= 2
	}

	for i := 3; i <= n; i++ {
		for n%i == 0 {
			faktor = append(faktor, i)
			n /= i
		}
	}

	return faktor
}
