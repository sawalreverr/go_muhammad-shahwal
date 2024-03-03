package main

import (
	"05-Data-Structure-Array-Slice-Map-Function/praktikum/prioritas-1/task"
	"fmt"
)

func main() {
	// task 1
	res1 := task.Merge([][]int{
		{1, 2, 5, 4, 3},
		{1, 2, 7, 8},
	})
	fmt.Println(res1) // [1,2,5,4,3,7,8]

	res2 := task.Merge([][]int{
		{1, 2, 5, 4, 5},
		{7, 9, 10, 5},
	})
	fmt.Println(res2) // [1, 2, 5, 4, 7, 9, 10]

	res3 := task.Merge([][]int{
		{1, 4, 5},
		{7},
	})
	fmt.Println(res3) // [1, 4, 5, 7]

	// task 2
	fmt.Println(task.PrimeSum([]int{1, 2, 4, 5, 12, 19, 30})) // 26
	fmt.Println(task.PrimeSum([]int{45, 17, 44, 67, 11}))     // 95
	fmt.Println(task.PrimeSum([]int{15, 12, 9, 10}))          // 0

	// task 3
	data1 := []float64{1, 1, 2, 5, 6, 8, 12, 4, 5, 5, 5, 8, 9}

	fmt.Printf("sum: %.2f\n", task.Sum(data1))       // 71.00
	fmt.Printf("mean: %.2f\n", task.Mean(data1))     // 5.46
	fmt.Printf("median: %.2f\n", task.Median(data1)) // 5.00
	fmt.Printf("mode: %.2f\n", task.Mode(data1))     // 5.00

	data2 := []float64{6, 7, 1, 11, 8, 12, 6, 12, 1, 7, 8, 10, 14}

	fmt.Printf("\nsum: %.2f\n", task.Sum(data2))  // 103.00
	fmt.Printf("sum: %.2f\n", task.Mean(data2))   // 7.92
	fmt.Printf("sum: %.2f\n", task.Median(data2)) // 8.00
	fmt.Printf("sum: %.2f\n", task.Mode(data2))   // 1.00
}
