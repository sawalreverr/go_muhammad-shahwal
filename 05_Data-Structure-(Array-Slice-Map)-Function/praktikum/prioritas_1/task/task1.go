package task

import "slices"

func Merge(data [][]int) []int {
	var slcNew []int

	for _, x := range data {
		for _, y := range x {
			if !slices.Contains(slcNew, y) {
				slcNew = append(slcNew, y)
			}
		}
	}

	return slcNew
}
