## Problem Solving Paradigm

### Problem Solving

Problem solving adalah kemampuan menyelesaikan masalah dengan pengambilan keputusan yang paling tepat.

Beberapa prinsipnya ialah:

- Brute Force (Complete Search)
- Divide & Conquer
- Greedy
- Dynamic Programming

### Brute Force

Brute Force (Complete Search) adalah metode dalam menyelesaikan sebuah masalah dengan menelusuri seluruh/setiap data untuk mendapatkan solusi yang dibutuhkan.

```go
// Contoh dari Brute Force (complete search) dalam pencarian nilai terbesar dan nilai terkecil
func FindMaxMin(numbers []int) (int, int) {
	max := numbers[0]
	min := numbers[0]

	for i := 0; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}

		if numbers[i] < min {
			min = numbers[i]
		}
	}

	return max, min
}

numbers := []int{10, 7, 3, 5, 8, 2, 9}
max, min := FindMaxMin(numbers) // 10, 2
```

### Divide and Conquer

Divide and Conquer adalah metode dalam menyelesaikan sebuah masalah dengan membaginya menjadi bagian-bagian yang lebih kecil dan kemudian menyelesaikan setiap bagian.

- Divide: Membagi masalah yang besar menjadi masalah yang lebih kecil
- Conquer: Ketika masalah sudah cukup kecil untuk diselesaikan, langsung selesaikan
- Combine: Jika dibutuhkan maka perlu

```go
// Contoh Divide & Conquer pada pencarian data, O(log n)
func BinarySearch(numbers []int, target int) int {
	hasil := 0
	kiri := 0
	kanan := len(numbers) - 1

	// ketika nilai kiri lebih kecil atau sama dengan kanan
	for kiri <= kanan {
		// jumlahkan nilai kiri + nilai kanan dibagi 2, untuk mendapatkan nilai tengahnya
		tengah := (kiri + kanan) / 2
		// jika target lebih kecil dari numbers nilai tengah yang sekarang
		if target < numbers[tengah] {
			// maka kanan akan bergeser sampai dengan tengah - 1,
			// contoh [1,1,3,5,5,7,8] maka nilai kanan yang sekarang adalah 3 pada index ke-2 dikarenakan tengah(5 atau index ke-3) - 1
			kanan = tengah - 1
			// jika target lebih besar dari numbers nilai tengah yang sekarang
		} else if target > numbers[tengah] {
			// maka kiri akan bergeser sampai dengan tengah + 1
			// contoh [1,1,3,5,5,7,8] maka nilai kiri yang sekarang adalah 5 pada index ke-5 dikarenakan tengah(5 atau index-ke) - 1
			kiri = tengah + 1
			// jika tidak maka kita sudah menemukan hasil == target nya secara langsung, dan lakukan break untuk menghentikan loop
		} else {
			hasil = tengah
			break
		}
	}

	return hasil
}

numbers := []int{1, 1, 3, 5, 5, 7, 8, 10}
BinarySearch(numbers, 3) // 2


// Contoh Divide & Conquer pada perpangkatan, O(log n)
func Pow(x, y int) int {
	if y == 0 {
		return 1
	}

	if y%2 == 0 {
		return Pow(x*x, y/2)
	} else {
		return x * Pow(x*x, y/2)
	}
}

Pow(2,8) // 256
```

### Greedy

Greedy adalah metode dalam penyelesaian sebuah masalah yang memilih opsi yang paling tepat berdasarkan situasi saat ini. Algoritma ini mengabaikan fakta bahwa hasil terbaik saat ini mungkin tidak memberikan hasil yang optimal secara keseluruhan.

```go
// Contoh dalam coinchange
func CoinChange(numbers []int, target int) []int {
	var change []int
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))

	for target != 0 {
		for _, num := range numbers {
			if target >= num {
				target -= num
				change = append(change, num)
			}
		}
	}

	return change
}

numbers := []int{10, 25, 5, 1}
CoinChange(numbers, 42) // [25 10 5 1 1]
```

### Dynamic Programming

Dynamic Programming adalah metode dalam penyelesaian sebuah masalah dengan cara memecahkan masalah yang kompleks dengan memecahnya menjadi submasalah yang lebih sederhana. Dengan menyelesaikan setiap submasalah hanya sekali dan menyimpannya hasilnya, metode ini menghindari komputasi yang berlebihan, sehingga menghasilkan solusi yang lebih efisien untuk berbagai masalah. Contoh masalah yang dapat kita selesaikan dengan Dynamic Programming ialah perhitungan Fibonacci seperti yang sudah saya buat sebelumnya.
