## String - Advance Function - Pointer - Method - Struct - Interface

### String

```go
import "strings"

var word string = "Kopi Kita"

// 1. Access String
word[0:]  // Kopi Kita
word[1:6] // opi K
word[:7]  // Kopi Ki

// 2. Compare String
word == "KopiKita" // false

// 3. Contain
strings.Contains(word, "Kopi") // true

// 4. Substring
wordKita := word[:5] // Kopi

// 5. Replace
strings.Replace(word, "i", "o", -1) // replace all char i to o, Kopo Kota

// 6. Insert
wordNew := word[:5] + "Kenangan " + word[5:] // Kopi Kenangan Kita
```

### Function

```go
// normal function
func sum(a, b int) int {}

// variadic function
func sum(data ...int) int {}
sum(1,2,3,4,5,6) // kita dapat memasukkan banyak int ke dalam parameter data

// anonymous function - 3 contoh
// contoh 1 : lgsung panggil
func(){
    fmt.Println("Hello World")
}()

// contoh 2 : assign to variable
anonFunc := func() {
    fmt.Println("Hello World")
}
anonFunc() // Hello World

// contoh 3 : with parameter
func(name string) {
    fmt.Println("Hello World " + name)
}("Budi")

// closure function
func NewCounter() func() int {
    count := 0
    return func() int {
        count += 1
        return count
    }
}
counter := NewCounter()
counter() // 1
counter() // 2 >> dia akan mereference count sebelumnya sehingga datanya akan bertambah

// defer function
defer func() {
    fmt.Println("Hello")
}()
fmt.Println("World")


// World
// Hello >> diakhir

// defer function itu seperti closer dan pasti akan dijalankan di akhir program
```

### Pointer ( \* | & )

Pointer adalah variabel yang menyimpan memory address dari variabel lain. Pointer memiliki kuasa untuk mengubah data dari variabel memorynya.

```go
word := "Budi"
newWord := &word

fmt.Println(word, &word)       // Budi 0xc000030070
fmt.Println(*newWord, newWord) // Budi 0xc000030070

*newWord = "Ilham"
fmt.Println(word, &word)       // Ilham 0xc000030070
fmt.Println(*newWord, newWord) // Ilham 0xc000030070
```

### Struct

Struct adalah kumpulan definisi variabel (atau property) dan atau fungsi (atau method), yang dibungkus sebagai tipe data baru dengan nama tertentu (mirip seperti OOP pada bahasa lain).

```go
// deklarasi
type Student struct {
    Name string
    FullName string
    Age int
    Address string
}

// use
// contoh 1
var s1 Student
s1.Name = "Ilham"
s1.FullName = "Ilham Kurniawan"
s1.Age = 17
s1.Address = "Jalan Kepiting Raya"
fmt.Println(s1) // {Ilham Ilham Kurniawan 17 Jalan Kepiting Raya}

// contoh 2
var s2 Student = Student{
    Name:     "Kurniawan",
    FullName: "Kurniawan Ilham",
    Age:      18,
    Address:  "Jalan Raya Kepiting",
}
fmt.Println(s2) // {Kurniawan Kurniawan Ilham 18 Jalan Raya Kepiting}

// contoh 3
s3 := Student{"Bagas", "Bagas Rizki", 19, "Jalan Bebek Sebrang"} // sesuai urutan
fmt.Println(s3) // {Bagas Bagas Rizki 19 Jalan Bebek Sebrang}
```

### Method

Method adalah fungsi yang menempel pada type (bisa struct atau tipe data lainnya). Method bisa diakses lewat variabel objek.

```go
type Student struct {
	Name     string
	FullName string
	Age      int
	Address  string
}

// method SayHello dari Student
func (s Student) SayHello() {
	fmt.Println("Hello ", s.FullName)
}

// method GetAllData dari Student
func (s Student) GetAllData() {
	fmt.Println("Nama panggilan: ", s.Name)
	fmt.Println("Nama lengkap: ", s.FullName)
	fmt.Println("Umur: ", s.Age)
	fmt.Println("Alamat: ", s.Address)
}

s := Student{
    Name:     "Ilham",
    FullName: "Ilham Kurniawan",
    Age:      17,
    Address:  "Jalan Durian Runtuh",
}

s.SayHello() // Hello Ilham Kurniawan

s.GetAllData()
// Nama panggilan:  Ilham
// Nama lengkap:  Ilham Kurniawan
// Umur:  17
// Alamat:  Jalan Durian Runtuh
```

### Interface

Interface adalah kumpulan definisi method yang tidak memiliki isi (hanya definisi saja), yang dibungkus dengan nama tertentu.

```go
// Deklarasi interface Hitung
type Hitung interface {
	Luas() float64
	Keliling() float64
}

// Struct Lingkaran
type Lingkaran struct {
	Diameter float64
}

// Struct Persegi
type Persegi struct {
	Sisi float64
}

// Method Lingkaran
func (l *Lingkaran) JariJari() float64 {
	return l.Diameter / 2
}
func (l *Lingkaran) Luas() float64 {
	return math.Pi * math.Pow(l.JariJari(), 2)
}
func (l *Lingkaran) Keliling() float64 {
	return math.Pi * l.Diameter
}

// Method Persegi
func (p *Persegi) Luas() float64 {
	return math.Pow(p.Sisi, 2)
}
func (p *Persegi) Keliling() float64 {
	return p.Sisi * 4
}


var bangunDatar Hitung

bangunDatar = &Persegi{10.0}
fmt.Println("Persegi")
fmt.Println("Luas: ", bangunDatar.Luas())
fmt.Println("Keliling: ", bangunDatar.Keliling())

bangunDatar = &Lingkaran{14.0}
fmt.Println("\nLingkaran")
fmt.Printf("Luas: %.2f\n", bangunDatar.Luas())
fmt.Printf("Keliling: %.2f\n", bangunDatar.Keliling())
fmt.Printf("Jari-jari: %.2f\n", bangunDatar.(*Lingkaran).JariJari()) // dikarenakan interface Hitung tidak memiliki method JariJari() jadi kita hrus buat external

// Persegi
// Luas:  100
// Keliling:  40

// Lingkaran
// Luas: 153.94
// Keliling: 43.98
// Jari-jari: 7.00
```
