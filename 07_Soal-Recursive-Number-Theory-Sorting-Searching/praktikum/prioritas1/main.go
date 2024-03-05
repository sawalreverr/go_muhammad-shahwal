package main

import (
	"fmt"
	"praktikum/prioritas1/task"
)

func main() {
	fmt.Println("Task 01")
	fmt.Println(task.FibX(5))  // 12
	fmt.Println(task.FibX(10)) // 143

	fmt.Println("\nTask 02")
	students := []task.Student{
		{"jamie", 67, 88, 90},
		{"michael", 77, 77, 90},
		{"aziz", 100, 65, 88},
		{"jamal", 50, 80, 75},
		{"eric", 70, 80, 65},
		{"john", 80, 76, 68},
	}
	task.GetInfo(students)

	fmt.Println("\nTask 03")
	fmt.Println(task.GetSequence(4))   // 10
	fmt.Println(task.GetSequence(15))  // 120
	fmt.Println(task.GetSequence(100)) // 5050
}
