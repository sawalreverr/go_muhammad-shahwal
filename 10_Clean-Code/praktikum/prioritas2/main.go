package main

import "fmt"

// ubah menjadi DataSet
type DataSet struct {
	Data map[int]int
}

func (s DataSet) Add(i int) {
	s.Data[i] = 1
}

func (s DataSet) Get() []int {
	r := []int{}

	for k := range s.Data {
		r = append(r, k)
	}

	return r
}

func (s DataSet) Remove(key int) {
	delete(s.Data, key)
}

func main() {
	dataSet := DataSet{
		Data: map[int]int{},
	}

	dataSet.Add(1)
	dataSet.Add(2)
	dataSet.Add(1)
	dataSet.Add(3)
	dataSet.Add(4)
	dataSet.Add(5)

	dataSet.Remove(100)

	fmt.Println(dataSet.Get())
}
