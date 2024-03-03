## Data Structure

### Array

Array adalah sebuah tipe data yang dapat menyimpan banyak data di dalam satu variabel. Pada array kita wajib mendefinisikan berapa banyak data yang ingin kita masukkan kedalam sebuah variabel (fix).

```go
var arr1 [5]int = [5]int{1,2,3,4,5} // sebuah array berisi 5 buah int
var arr2 [5]string = [5]string{"a","b","c","d","e"} // sebuah array berisi 5 buah string
arr3 := [5]int{1,2,3} // kita juga bisa membuat seperti ini, jika kita print maka [1 2 3 0 0]
```

### Slice

Slice adalah sebuah tipe data yang dapat menyimpan banyak data di dalam satu variabel. Tidak seperti array pada slice kita tidak perlu mendefinisikan jumlah capacity dari slice tersebut (dynamic).

```go
var slc1 []int = []int{1,2,3}
var slc2 []string = []int{"red","blue","green"}

len(slc1) // 3
cap(slc1) // 3

append(slc1, 4) // menambah nilai 4 kedalam slc1
cap(slc1) // 6, kenapa capacitynya berubah menjadi 6? hal ini dikarenakan ketika kita menambah/append pada slice maka capacity dari slice sebelumnya akan dikalikan 2 sehingga yang sebelumnya 3 menjadi 6, sampai data tersebut berjumlah 6 baru akan dikalikan 2 lgi
```

### Map

Map adalah sebuah tipe data yang dapat menyimpan banyak data di dalam satu variabel. Pada tipe data Map terdapat Key dan Value yaitu sebuah fungsi untuk menetapkan sebuah nilai, sebuah Key bersifat unik (berbeda2).

```go
// map[typeDataKey]typeDataValue = Key(fix) - Value(dynamic)

hargaGame := map[string]int{"Elden Ring": 600000, "Palworld": 250000}
fmt.Println(hargaGame) // map[Elden Ring:600000 Palworld:250000]

// mengubah value pada map
hargaGame["Palworld"] = 260000 // mengubah data dengan key Palworld menjadi 260000
fmt.Println(hargaGame) // map[Elden Ring:600000 Palworld:260000]

// membuat map dengan make()

hargaBuah := make(map[string]int)
hargaBuah["Apel"] = 20000
hargaBuah["Jeruk"] = 15000
hargaBuah["Nanas"] = 12000

fmt.Println(hargaBuah) // map[Apel:20000 Jeruk:15000 Nanas:12000]
```

### Function

Function adalah sebuah blok kode yang dirancang untuk melakukan tugas tertentu dan dapat dipanggil berkali-kali dalam program.

```go
// deklarasi function : func namaFunction() {isi kode}

// contoh function
func SayHello() {
    fmt.Println("Hello")
}

// contoh function dengan parameter
func SayHelloWithName(name string) {
    fmt.Println("Hello " + name)
}

// contoh function dengan parameter dan pengembalian
func SayHelloWithNameReturn(name string) string {
    return "Hello " + name
}

// contoh function dengan parameter slice int dan pengembalian
func Sum(arr ...int) int {
	var total int
	for _, num := range arr {
		total += num
	}

	return total
}
```

## Packages

Banyak packages yang sudah disediakan oleh golang itu sendiri, sebagai contoh pada update go 1.21 packages slices dirilis oleh golang.

### Packages Slice

#### Slices

[pkg.go.dev/slices](https://pkg.go.dev/slices)

```go
import "slices"

// mencari angka 2 didalam sebuah slice dengan BinarySearch
slices.BinarySearch[[]int](arr, 2)

// mencari people yang bernama "Budi" di dalam sebuah slice Person
slices.BinarySearchFunc[[]Person](people, Person{"Budi", 0}, func(a, b Person) int {
	return cmp.Compare(a.Name, b.Name)
})

// menghapus duplicate yang ada pada slice
slc := []int{0, 1, 1, 2, 3, 5, 8}
slices.Compact(slc) // [0 1 2 3 5 8]

// mengecek keberadaan elemen pada slice
slices.Contains[[]string](arr, name)

// menghapus item berdasarkan index pada slice
slice.Delete(arr, 1, 4) // start(1) until(4) = (index 1 - index 3)

// menghapus item berdasarkan kondisi pada slice
slices.DeleteFunc[[]int](arr, func(number int) bool {
	return number % 2 == 0 // menghapus semua data yang genap
})

// mencari apple pada index ke berapa pada slice
slices.Index[[]string](arr, "apple")

// insert nilai ke dalam sebuah slice
slices.Insert[[]int](arr, 1, 20) // insert nilai 20 ke index 1

// mengecek apakah slice itu berurutan atau tidak
slices.IsSorted[[]int](arr)

// mengambil nilai terbesar dan juga nilai terkecil pada sebuah slice
slices.Max[[]int](arr) // untuk nilai terbesar
slices.Min[[]int](arr) // untuk nilai terkecil

// mereplace nilai/beberapa nilai berdasarkan index pada sebuah slice
slices.Replace[[]string](arr, 0, 3, "g","h","j") // replace index 0 - 3-1, dengan nilai g h j

// reverse sebuah slice
slices.Reverse[[]string](arr)

// sorting sebuah slice
slices.Sort[[]int](arr)

// sorting sebuah slice dengan kondisi pada slice
slices.SortFunc[[]Product](products, func(p1, p2 Product) int {
    return cmp.Compare[int](p1.Price, p2.Price) // akan mensorting berdasarkan harga/price
})
```

#### maps

[pkg.go.dev/maps](https://pkg.go.dev/maps)

```go
import "maps"

students := map[string]int{
    "Ilham": 80,
    "Kurniawan": 75,
    "Santoso": 85,
    "Rizki": 90,
}

// menghapus nilai berdasarkan kondisi
maps.DeleteFunc[map[string]int](students, func(name string, score int) bool {
    return score < 80
})

// output: map["Ilham": 80 "Santoso": 85 "Rizki": 90]
```
