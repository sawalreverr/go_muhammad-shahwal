package main

import (
	"05-Data-Structure-Array-Slice-Map-Function/praktikum/prioritas-2/task"
	"fmt"
)

func main() {
	// task 1
	fmt.Println(task.SpinSlice([]int{1, 2, 3, 4, 5}))    // [1,5,2,4,3]
	fmt.Println(task.SpinSlice([]int{6, 7, 8}))          // [6,8,7]
	fmt.Println(task.SpinSlice([]int{8, 5, 4, 3, 2, 1})) // [8,1,5,2,4,3]

	// task 2
	fmt.Println(task.GroupPrime([]int{2, 3, 4, 5, 6, 7, 8})) // [[2,3,5,7],4,6,8]
	fmt.Println(task.GroupPrime([]int{15, 24, 3, 9, 5}))     // [[3,5],15,24,9]
	fmt.Println(task.GroupPrime([]int{4, 8, 9, 12}))         // [4,8,9,12]
}
