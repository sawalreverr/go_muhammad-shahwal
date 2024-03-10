package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Task 01")
	fmt.Println(ArrayAns(2))
	fmt.Println(ArrayAns(3))

	fmt.Println("\nTask 02")
	fmt.Println(SegitigaPascal(3))
	fmt.Println(SegitigaPascal(5))
	fmt.Println(SegitigaPascal(7))

	fmt.Println("\nTask 03")
	fmt.Println(Fibonacci(0))
	fmt.Println(Fibonacci(2))
	fmt.Println(Fibonacci(9))
	fmt.Println(Fibonacci(10))
	fmt.Println(Fibonacci(12))

	fmt.Println("\nTask 04")
	SimpleEquations(1, 2, 3)
	SimpleEquations(6, 6, 14)
	SimpleEquations(9, 24, 29)
}

func ArrayAns(n int) []string {
	var ans []string
	for i := 0; i <= n; i++ {
		ans = append(ans, strconv.FormatInt(int64(i), 2))
	}
	return ans
}

func SegitigaPascal(n int) [][]int {
	var res [][]int

	for i := 0; i < n; i++ {
		base := 1
		tmpSlc := []int{}
		for j := 0; j <= i; j++ {
			if !(i == 0 || j == 0) {
				base = base * (i - j + 1) / j
			}
			tmpSlc = append(tmpSlc, base)
		}
		res = append(res, tmpSlc)
	}

	return res
}

var cacheFib = map[int]int{}
var isFound = false

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	fib1, isFound := cacheFib[n-1]
	if !isFound {
		fib1 = Fibonacci(n - 1)
		cacheFib[n-1] = fib1
	}

	fib2, isFound := cacheFib[n-2]
	if !isFound {
		fib2 = Fibonacci(n - 2)
		cacheFib[n-2] = Fibonacci(n - 2)
	}

	return fib1 + fib2
}

func SimpleEquations(a, b, c int) {
	var x, y, z int

	for x+y+z != a {
		if y <= x {
			y++
		} else if z <= y {
			z++
		} else {
			x++
			y++
			z++
		}
	}

	if (x*y*z == b) && (x*x+y*y+z*z == c) {
		fmt.Println(x, y, z)
	} else {
		fmt.Println("No Solution")
	}
}
