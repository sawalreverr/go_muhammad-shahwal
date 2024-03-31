# ORM & Code Structure (MVC)

## What is Object Relational Mapping (ORM)

ORM (Object Relation Mapping) merupakan teknik yang merubah suatu table menjadi sebuah object yang nantinya mudah untuk digunakan. Object yang dibuat memiliki property yang sama dengan field-field yang ada pada table tersebut.

ORM memungkinkan kita melakukan query dan memanipulasi data di database menggunakan object oriented.

## What is GORM

GORM adalah salah satu ORM (Object-Relational Mapping) yang populer digunakan dalam bahasa pemrograman Go (Golang).

Dengan GORM, kita dapat mempermudah pengoperasian basis data relasional, seperti menghubungkan, memperoleh, menyimpan, dan mengubah data dengan mudah.

```bash
go get -u gorm.io/gorm
```

## Database Connection & Migration

Auto migration di GORM adalah fitur yang memungkinkan developer untuk secara otomatis membuat dan mengelola skema database dari definisi model GORM mereka. Dengan menggunakan fitur ini, pengembang dapat menghindari kebutuhan untuk secara manual membuat skema database dan mengubahnya ketika ada perubahan dalam model aplikasi mereka.

Untuk menggunakan auto migration di GORM, kita dapat memanggil metode **AutoMigrate** pada objek **DB** Anda, yang akan membuat tabel untuk setiap model GORM dan kolom untuk setiap bidang di model. Jika tabel atau kolom sudah ada di database, GORM akan memperbarui struktur tersebut sesuai dengan definisi model terbaru.

Contoh penggunaan auto migration di GORM dapat dilihat di bawah ini:

```go
import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name string
    Email string
}

func main() {
    dsn := "user:password@tcp(localhost:3306)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    db.AutoMigrate(&User{})
}
```

Dalam contoh ini, kita membuat struktur User sebagai model GORM dan kemudian memanggil metode AutoMigrate untuk membuat tabel users di database dan kolom untuk setiap bidang di model User. Jika kita menambahkan atau mengubah bidang di model User, kita dapat menjalankan ulang aplikasi untuk memperbarui skema database.

Untuk lebih lanjut, silahkan kunjungi docs gormnya langsung [GORM](https://gorm.io/docs/)

## What is MVC

Model View Controller atau yang dapat disingkat MVC adalah sebuah pola arsitektur dalam membuat sebuah aplikasi dengan cara memisahkan kode menjadi tiga bagian yang terdiri dari:

- Model\
  Bagian yang bertugas untuk menyiapkan, mengatur, memanipulasi, dan mengorganisasikan data yang ada di database.

- View\
  Bagian yang bertugas untuk menampilkan informasi dalam bentuk Graphical User Interface (GUI).

- Controller\
  Bagian yang bertugas untuk menghubungkan serta mengatur model dan view agar dapat saling terhubung.

Berikut ini adalah alur prosesnya.

- Proses pertama adalah view akan meminta data untuk ditampilkan dalam bentuk grafis kepada pengguna.
- Permintaan tersebut diterima oleh controller dan diteruskan ke model untuk diproses.
- Model akan mencari dan mengolah data yang diminta di dalam database
- Setelah data ditemukan dan diolah, model akan mengirimkan data tersebut kepada controller untuk ditampilkan di view.
- Controller akan mengambil data hasil pengolahan model dan mengaturnya di bagian view untuk ditampilkan kepada pengguna.

## Why Need Structure

Ada beragam manfaat ketika kita menerapkan MVC ini dalam pembuatan aplikasi kita. Berikut ini adalah manfaatnya.

- Proses pengembangan aplikasi menjadi lebih efisien\
  Penggunaan MVC dapat mempercepat pengembangan aplikasi karena kode dapat dikerjakan oleh beberapa developer. Contohnya dalam kasus pengembangan aplikasi web, bagian model dan controller dapat dikerjakan oleh back-end developer sedangkan bagian view dapat dilakukan oleh front-end developer.

- Penulisan kode menjadi lebih rapi\
  Karena dibagi menjadi tiga bagian, maka penulisan kode akan jadi lebih rapi dan memudahkan developer lain untuk mengembangkan kode tersebut.

- Dapat melakukan testing dengan lebih mudah\
  Untuk memastikan seluruh aplikasi bekerja sesuai dengan rencana maka langkah testing atau uji coba wajib dilakukan. Dengan menggunakan model view controller ini, maka proses uji coba dapat dilakukan pada setiap bagian.

- Perbaikan bug atau error lebih cepat untuk diselesaikan\
  Penggunaan MVC dapat memudahkan developer untuk memperbaiki error atau bug yang terjadi. Developer dapat fokus untuk menemukan dan memperbaiki masalah yang terjadi karena kode dituliskan pada bagian-bagian terpisah.

- Mempermudah pemeliharaan\
  Konsep MVC ini dapat mempermudah pemeliharaan aplikasi, karena script atau kode yang lebih rapi dan terstruktur sehingga mempermudah developer dalam proses pemeliharaan aplikasi.

## Structuring Project

Ada tiga komponen atau bagian MVC yaitu model, view, dan controller.
Berikut penjelasan dari masing-masing komponen MVC dan fungsinya:

- Model\
  Komponen model berhubungan dengan seluruh data logis (data-related logic) yang digunakan pengguna.
  Data logis adalah data yang relevan dengan konteks tertentu. Misalnya ketika ingin membuat database pelanggan, maka data yang dikumpulkan berupa nama pelanggan, usia, alamat, nomor kontak, dan riwayat pembelian.
  Komponen model dapat mewakili data apa saja yang sedang ditransfer antara komponen view dan controller, atau data logis lainnya.
  Selain itu, model juga dapat mengambil data dari database, memodifikasi, dan menambahkan data ke dalam database.

- View\
  Komponen view bertanggung jawab untuk membuat tampilan muka (UI/user interface) pada seluruh situs atau aplikasi.
  Komponen ini dibuat dari data yang dikumpulkan oleh model, dan diberikan kepada view melalui komponen controller. Namun pada beberapa kasus, view bisa berinteraksi langsung dengan model tanpa bantuan controller.

- Controller\
  Komponen controller berfungsi sebagai penghubung antara model dengan view.
  Tugas controller hanyalah memproses data dan permintaan yang masuk, kemudian memberitahu komponen model apa yang harus dikerjakan, dan hasilnya akan ditampilkan oleh komponen view.
