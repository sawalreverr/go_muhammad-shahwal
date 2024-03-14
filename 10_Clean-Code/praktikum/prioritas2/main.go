package main

import "fmt"

// ubah menjadi DataSet
type DataSet struct {
	Data map[int]int
}

func (ds DataSet) Add(i int) {
	ds.Data[i] = 1
}

func (ds DataSet) Get() []int {
	r := []int{}

	for k := range ds.Data {
		r = append(r, k)
	}

	return r
}

func (ds DataSet) Remove(key int) {
	delete(ds.Data, key)
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
