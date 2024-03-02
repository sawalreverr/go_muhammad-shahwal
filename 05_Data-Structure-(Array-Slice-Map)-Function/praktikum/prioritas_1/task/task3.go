package task

import (
	"math"
	"slices"
)

func Sum(data []float64) float64 {
	var sum float64

	for _, v := range data {
		sum += v
	}

	return (math.Floor(sum*100) / 100)
}

func Mean(data []float64) float64 {
	mean := Sum(data) / float64(len(data))

	return (math.Floor(mean*100) / 100)
}

func Median(data []float64) float64 {
	slices.Sort(data)
	mid := len(data) / 2

	if len(data)%2 != 0 {
		return data[mid]
	} else {
		return (data[mid-1] + data[mid]) / 2.0
	}
}

func Mode(data []float64) float64 {
	dataset := make(map[float64]float64)
	var max float64
	var mode float64

	for _, d := range data {
		dataset[d] += 1
	}

	for k, v := range dataset {
		if v > max {
			max = v
			mode = k
		}
	}

	return mode
}
