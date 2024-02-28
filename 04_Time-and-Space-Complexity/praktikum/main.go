package main

import (
	"fmt"
	"math"
)

func main() {
	// Task 1
	fmt.Println(primeNumber(1000000007))
	fmt.Println(primeNumber(13))
	fmt.Println(primeNumber(17))
	fmt.Println(primeNumber(20))
	fmt.Println(primeNumber(35))

	// Task 2
	fmt.Println(pow(2, 3))
	fmt.Println(pow(5, 3))
	fmt.Println(pow(10, 2))
	fmt.Println(pow(2, 5))
	fmt.Println(pow(7, 3))
}

func primeNumber(n int) bool {
	if n <= 1 { // menghindari nilai kurang dari 2
		return false
	} else if n == 2 { // jika inputan 2 maka lgsung saja
		return true
	} else if n%2 == 0 { // lgsung mengembalikan jika n dapat dibagi 2, jdi tidak perlu masuk ke looping
		return false
	}

	x := int(math.Sqrt(float64(n)))
	for i := 3; i <= x; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true

	// mungkin ini bisa disebut dengan O(log n)
	// referensi, dengan sedikit perubahan https://www.linkedin.com/pulse/optimized-prime-number-algorithm-surya-satish-vegisetti/
}

func pow(x, n int) int {
	if n == 0 {
		return 1
	} else if n%2 == 1 { // jika ganjil
		return x * pow(x, n-1)
	} else {
		tmp := pow(x, n/2)
		return tmp * tmp
	}

	// O(log n)
	// referensi, https://medium.com/@fazry/cara-rekursif-di-golang-cc202f55e336
}
