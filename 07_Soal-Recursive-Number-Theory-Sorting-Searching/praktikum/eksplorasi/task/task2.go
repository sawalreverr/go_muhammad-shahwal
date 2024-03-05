package task

import "sort"

func GetItem(items []int, target int) []int {
	var output []int
	sort.Sort(sort.Reverse(sort.IntSlice(items)))

	for _, item := range items {
		if target != 0 {
			target -= item
			output = append(output, item)
		} else {
			break
		}

	}

	return output
}
