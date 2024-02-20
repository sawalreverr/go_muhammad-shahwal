# Basic Programming

### Introduction Golang

#### Apa itu Golang ?

Golang (atau biasa disebut dengan Go) adalah bahasa pemrograman yang dikembangkan di Google oleh Robert Griesemer, Rob Pike, dan Ken Thompson pada tahun 2007 dan mulai diperkenalkan ke publik tahun 2009.

Penciptaan bahasa Go didasari bahasa C dan C++, oleh karena itu gaya sintaksnya mirip.

#### Kelebihan Golang
Go memiliki kelebihan dibanding bahasa lainnya, beberapa di antaranya:
- Mendukung konkurensi di level bahasa dengan pengaplikasian cukup mudah
- Mendukung pemrosesan data dengan banyak prosesor dalam waktu yang bersamaan (pararel processing)
- Memiliki garbage collector
- Proses kompilasi sangat cepat
- Bukan bahasa pemrograman yang hirarkial dan bukan strict OOP, memberikan kebebasan ke developer perihal bagaimana cara penulisan kode.
- Dependensi dan tooling yang disediakan terbilang lengkap.
- Dukungan komunitas sangat bagus. Banyak tools yang tersedia secara gratis dan open source yang bisa langsung dimanfaatkan.

### Installation Golang

#### Instalasi Go di Windows

- Download terlebih dahulu installer-nya di https://go.dev/dl/. Pilih installer untuk sistem operasi Windows sesuai jenis bit yang digunakan.
- Setelah ter-download, jalankan installer, klik next hingga proses instalasi selesai. By default jika anda tidak merubah path pada saat instalasi, Go akan ter-install di C:\go. Path tersebut secara otomatis akan didaftarkan dalam PATH environment variable.
- Buka Command Prompt / CMD, eksekusi perintah berikut untuk mengecek versi Go.
```
go version
```
- Jika output adalah sama dengan versi Go yang ter-install, menandakan proses instalasi berhasil.

#### Instalasi Go di Linux

- Unduh arsip installer dari https://go.dev/dl/, pilih installer untuk Linux yang sesuai dengan jenis bit komputer anda. Proses download bisa dilakukan lewat CLI, menggunakan wget atau curl.
```
$ wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
```
- Buka terminal, extract arsip tersebut ke /usr/local.
```
$ rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
```
- Tambahkan path binary Go ke PATH environment variable.
```
$ echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc
$ source ~/.bashrc
```
- Selanjutnya, eksekusi perintah berikut untuk mengetes apakah Go sudah terinstal dengan benar.
```
$ go version
```
- Jika output adalah sama dengan versi Go yang ter-install, menandakan proses instalasi berhasil.

### Mulai mencoba Golang 

Buka IDE favorit kamu, sebagai contoh saya menggunakan VS-Code sebagai IDE andalan saya. Kemudian buat folder workspace baru kamu dan buat file main.go di dalamnya. Tapi sebelum itu mari kita inisiasi terlebih dahulu. Buka terminal/cmd dan ketik dibawah ini: 

```
$ go mod init my_first_app
```

Perintah diatas berfungsi sebagai inisiasi go.mod pada workspace kamu.

#### Contoh Program
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello Master, This is Golang Languange")
}
```

```
$ go run main.go
Hello Master, This is Golang Languange
```

#### Command Go Terminal
| Command    | Function                                      |
| :--------- | :-------------------------------------------  |
| `$ go run` | Execute the program with extension **.go**    |
| `$ go build` | Compile program files |
| `$ go install` | Like go build and continue the install process |
| `$ go test` | For testing with suffix **_test.go** |
| `$ go get` | Download the Go package |

----

#### Apa itu Variable ?
Variable adalah value yang bisa berubah, berdasarkan kondisi atau informasi yang diberikan pada program. Biasanya variable ditandai dengan memberi nama dan tipe datanya (boolean, string, int, etc)

#### Type data di Golang

- Boolean (true, false)
- Integer (uint8, uint16, uint32, uint64, int8, int16, int32, int64, int, rune, uint, byte)
- Float (float32, float64)
- Complex (complex64, complex128)
- String 

#### Operators di Golang
- Arithmetic (+, -, *, /, %, ++, --)
- Comparison (==, !=, >, <, >=, <=)
- Logical (&&, ||, !)
- Bitwise (&, |, ^, <<, >>)
- Assignment (=, +=, -=, *=, /=, %= <<=, =>>, &=, ^=, |=)
- Miscellaneous (*, &) Pointer

#### IF Statement

```go
var myAge int = 17

if myAge == 17 {
    ...
} else if myAge < 12 {
    ...
} else {
    ...
}
```

#### Switch Statement

```go
var day int = 2

switch day {
    case day == 1:
        ...
    case day == 2:
        ...
    case day == 3:
        ...
    default:
        ...
} 
```

#### For Statement
```go
// Contoh 1
for i := 0; i < range; i++ {
    ...
}

// Contoh 2
sum := 0
for sum < 10 {
    sum += sum
}

// Contoh 3
for key, value := range variable {
    ...
}
```

