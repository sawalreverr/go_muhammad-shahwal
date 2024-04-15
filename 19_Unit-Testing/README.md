# Unit Testing

## What is Software Testing

Software Testing adalah sebuah proses untuk menganalisa sebuah sistem/software untuk mendeteksi kondisi fitur sistem saat ini dan fitur yang diinginkan oleh user/dev dan untuk mengetahui apakah fitur yang kita kerjakan sudah berjalan dengan baik.

## Purpose of Testing

- Preventing Regression

  Kondisi dimana software yang sebelumnya sudh berjalan dengan baik menjadi error dikarenakan adanya penambahan fitur

- Confidence Level in Refactoring

  Sebuah proses dimana kita mengubah sistem tanpa mengubah functionalitasnya

- Improve Code Design

  Semakin sering dev melakukan unit testing, maka dev akan memikirkan bagaimana kode yang dia baru buat tersebut dapat berjalan dengan baik hanya dengan memikirkannya

- Code Documentation

  Dengan adanya testing, kita akan mengetahui bagaimana proses sistem kita bekerja seperti bagaimana inputnya dan bagaimana outputnya

- Scaling

  Orang lain akan lebih mudah memahami kode kita, dengan adanya unit testing yang telah kita buat

## Level of Testing

- UI Test (End To End)

  Seperti namanya UI Test yaitu testing untuk tampilan aplikasi, sebagai contoh percobaan untuk register/login pada e-commerce

- Integration Test

  Integration test adalah testing untuk spesifik module tertentu, sebagai contoh dev melakukan testing pada API-nya apakah berjalan dengan baik atau tidak

- Unit Test

  Melakukan testing terhadap unit terkecil dalam sebuah aplikasi dalam membuat tugas sederhana

## Framework

Framework pada unit test adalah sebuah library dalam melakukan unit testing pada project kita, tanpa harus kita buat terlebih dahulu jdi tinggal kita gunakan.

Pada golang kita dapat menggunakan banyak sekali library unit test yang sudah tersedia, pada kali ini kita akan menggunakan library testify
https://github.com/stretchr/testify

## Structure

Pattern yang sering digunakan pada unit testing adalah

- Centralize (meletakkan semua file test kita ke dalam folder tests)
- Together (meletakkan setiap file test kita tepat disamping/didalam file production kita)

## Runner

Sebuah aplikasi/tool yang di desain untuk menjalankan sebuah test, contoh seperti bahasa go sudah tersedia tool runner testnya yaitu go test

## Mocking

Membuat object untuk mensimulasikan behaviour dari dependency yang digunakan itulah yang disebut mocking.

Jadi misal kita membuat fungsi untuk melakukan pengiriman email, tentu kita tidak ingin tiap test dijalankan email juga selalu dikirimkan. Karena kita tidak ingin menerima spam email dan juga test tidak boleh bergantung pada external services (Email service).

Dengan mocking kita dapat menghindari hal tersebut dengan membuat simulasi object yang sebenarnya.

### You Do Not To Test

- 3rd party modules
- 3rd party api or other external system
- Database
- Object class that you have test in other place
