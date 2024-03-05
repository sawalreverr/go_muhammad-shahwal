package main

import (
	"fmt"
	"praktikum/eksplorasi/task"
)

func main() {
	fmt.Println("Task 01")
	fmt.Println(task.GetDiamondSeq(4))   // 16
	fmt.Println(task.GetDiamondSeq(25))  // 625
	fmt.Println(task.GetDiamondSeq(100)) // 10000

	fmt.Println("\nTask 02")
	fmt.Println(task.GetItem([]int{10, 10, 5, 30, 40, 10, 5}, 80)) // [40 30 10]
	fmt.Println(task.GetItem([]int{50, 20, 10, 25, 25}, 100))      // [50 25 25]
}
