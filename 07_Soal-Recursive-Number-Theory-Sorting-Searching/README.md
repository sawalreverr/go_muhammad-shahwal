## Recursive - Number Theory - Sorting - Searching

### Recursive

Recursive/Rekursif adalah sebuah fungsi untuk menyelesaikan sebuah masalah dengan cara memanggil dirinya sendiri secara berulang.

```go
func factorial(n int) int {
    // !1
    if n == 1 {
        return 1
    }

    return n * factorial(n - 1) // 5! 4! 3! 2!
}

factorial(5) // 120
```

### Number Theory

Number Theory adalah cabang dari matematika murni yang dikhususkan terutama untuk mempelajari bilangan bulat dan fungsi aritmatika, sebagai contoh pada perhitungan Faktorial, Prima, FPB (Faktor Persekutuan Terbesar), KPK (Kelipatan Persekutuan Terkecil)

```go
func fpb(x,y int) int {
    if y == 0 {
        return x
    }

    return fpb(y, x%y)
}

func kpk(x,y int) int {
    return x * (y / fpb(x,y))
}

fpb(20, 45) // 5
fpb(9, 18)  // 9

kpk(12, 24) // 24
kpk(5, 20)  // 20
```

### Searching

Searching adalah sebuah algoritma untuk pencarian data dengan kriteria tertentu. Ada banyak sekali jenis2 algoritma untuk searching, seperti linear search, binary search, dll

```go
import (
    "fmt"
    "slices"
)

numbers := []int{4, 2, 7, 9, 1, 10}

// contoh linear search
func LinearSearch(numbers []int, target int) int {
	for i, num := range numbers {
		if num == target {
			return i
		}
	}
	return -1
}

LinearSearch(numbers, 7) // 2

// contoh built-in search
idx, found := slices.BinarySearch(numbers, 7)
fmt.Println(idx, found) // 2 true

```

### Sorting

Sorting adalah sebuah algoritma untuk menyusun sekumpulan data berdasarkan urutannya, baik berdasarkan terkecil maupun terbesar.

[Visualisasi Sorting](https://visualgo.net/en/sorting)

```go
// Selection Sort O(n^2)
func SelectionSort(numbers []int) []int {
	n := len(numbers)

	for i := 0; i < n; i++ {
		minimal := i
		for j := i + 1; j < n; j++ {
			if numbers[j] < numbers[minimal] {
				minimal = j
			}
		}
		numbers[minimal], numbers[i] = numbers[i], numbers[minimal]
	}

	return numbers
}

numbers1 := []int{4, 2, 7, 9, 1, 10}
SelectionSort(numbers1) // [1 2 4 7 9 10]

// Counting Sort O(n + k)
func CountingSort(numbers []int, k int) []int {
	count := make([]int, k+1)
	for i := 0; i < len(numbers); i++ {
		count[numbers[i]]++
	}

	counter := 0
	for i := 0; i < k+1; i++ {
		for j := 0; j < count[i]; j++ {
			numbers[counter] = i
			counter += 1
		}
	}

	return numbers
}

numbers2 := []int{7, 7, 7, 7, 1, 1, 1, 6, 6, 5, 5, 1, 2}
CountingSort(numbers2, len(numbers2)) // [1 1 1 1 2 5 5 6 6 7 7 7 7]

// Merge Sort O(log n)
func MergeSort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	middle := len(numbers) / 2
	left := numbers[:middle]
	right := numbers[middle:]
	numbers = Merge(MergeSort(left), MergeSort(right))
	return numbers
}

func Merge(left, right []int) []int {
	var i, j, k int
	result := make([]int, len(left)+len(right))

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}

	return result
}

numbers := []int{4, 2, 7, 9, 1, 10, 2}
MergeSort(numbers) // [1 2 2 4 7 9 10]
```

### Built-in sort

```go
import "sort"

// Sorting Int
numbers := []int{4, 2, 7, 9, 1, 10, 2}
sort.Ints(numbers)
fmt.Println(numbers) // [1 2 2 4 7 9 10]

// Sorting String
words := []string{"jeruk", "apel", "nanas", "pepaya"}
sort.Strings(words)
fmt.Println(words) // [apel jeruk nanas pepaya]

// Sorting Slices
products := []Product{
	{Name: "Apel", Price: 20},
	{Name: "Nanas", Price: 10},
	{Name: "Pepaya", Price: 30},
	{Name: "Jeruk", Price: 15},
	{Name: "Anggur", Price: 35},
}

// Berdasarkan Harga Terkecil
sort.Slice(products, func(i, j int) bool {
    return products[i].Price < products[j].Price
}) // [{Nanas 10} {Jeruk 15} {Apel 20} {Pepaya 30} {Anggur 35}]

// Berdasarkan Nama barang dan mulai dari awal abjad
sort.Slice(products, func(i, j int) bool {
    return products[i].Name < products[j].Name
}) // [{Anggur 35} {Apel 20} {Jeruk 15} {Nanas 10} {Pepaya 30}]
```
