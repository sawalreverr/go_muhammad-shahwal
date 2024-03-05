package main

import (
	"fmt"
	"praktikum/prioritas2/task"
)

func main() {
	fmt.Println("Task 01")
	fmt.Println(task.Catalan(7))  // 429
	fmt.Println(task.Catalan(10)) // 16796

	fmt.Println("\nTask 02")
	fmt.Println(task.PrimeFactors(20))   // 2, 2, 5
	fmt.Println(task.PrimeFactors(75))   // 3, 5, 5
	fmt.Println(task.PrimeFactors(1024)) // 2, 2, 2, 2, 2, 2, 2, 2, 2, 2
}
