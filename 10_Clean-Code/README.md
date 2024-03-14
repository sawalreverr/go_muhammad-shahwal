## Clean Code

### Apa itu Clean Code ?

Clean Code adalah istilah untuk kode yang mudah dibaca, dipahami dan diubah oleh programmer.

**_"Any fool can write code that a computer can understand. Good programmers write code that humans can understand"_** \
~ Martin Fowler

### Kenapa Clean Code ?

Work Collaboration, agar kode kita dapat dibaca oleh rekan tim kita. \
Feature Development, agar ketika ada penambahan fitur teman/rekan tim kita sudah langsung paham. \
Faster Development, agar kita tidak perlu ribet-ribet baca dari awal.

### Karakteristik Clean Code

- Penamaan yang mudah dipahami
- Mudah dieja dan dicari
- Singkat namun mendeskripsikan konteks
- Konsisten (jika camelcase maka seterusnya akan begitu)
- Hindari penambahan konteks yang tidak perlu
- Komentar (tidak setiap kode harus kita beri komentar)
- Function (tidak terlalu banyak argumen)
- Gunakan konvensi (style guide pada perusahaan besar seperti google atau airbnb)
- Formatting (pastikan kode kita rapih dan terhindar dari text kepanjangan maupun kelebaran, saran menggunakan prettier atau formatter lainnya)

### Clean Code Principle

#### KISS (Keep It So Simple)

Hindari membuat fungsi yang dibuat untuk melakukan A, sekaligus memodifikasi B, mengecek fungsi C, dst

Tips KISS

- Fungsi atau class harus kecil
- Fungsi dibuat untuk melakukan satu tugas saja
- Jangan gunakan terlalu banyak argumen pada fungsi
- Harus diperhatikan untuk mencapai kondisi yang seimbang, kecil dan jumlahnya minimal

#### DRY (Don't Repeat Yourself)

Duplikasi kode terjadi karena sering copy paste. Untuk menghindari duplikasi kode buatlah fungsi yang dapat digunakan secara berulang-ulang

#### Refactoring

Refactoring adalah proses restrukturisasi kode yang dibuat, dengan cara mengubah struktur internal tanpa mengubah perilaku eksternal. Prinsip KISS dan DRY bisa dicapai dengan cara refaktor.

Teknik Refactoring

- Membuat sebuah abstraksi
- Memecah kode dengan fungsi/class
- Perbaiki penamaan dan lokasi kode
- Deteksi kode yang memiliki duplikasi
