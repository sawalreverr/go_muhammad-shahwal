# Clean Architecture

Terdapat dua hal yang sering diketahui oleh developer dalam melakukan pengembangan suatu perangkat lunak, yaitu desain dan arsitektur. Kedua aspek tersebut saling berkaitan satu sama lain.

Dalam penjelasan yang disampaikan oleh Robert C. Martin dalam bukunya yang berjudul “Clean Architecture: A Craftsman’s Guide to Software Structure and Design”, tujuan dari desain dan arsitektur dalam perangkat lunak dapat diuraikan sebagai berikut.

“Meminimalkan sumber daya manusia yang dibutuhkan untuk membangun dan memelihara sistem yang dibutuhkan.
Ukuran kualitas desain hanyalah ukuran upaya yang diperlukan untuk memenuhi kebutuhan pelanggan. Jika upaya itu rendah, dan tetap rendah sepanjang masa pakai sistem, desainnya bagus. Jika upaya tersebut meningkat dengan setiap rilis baru, desainnya buruk.”

Arsitektur merupakan struktur dari suatu perangkat lunak yang dikembangkan oleh para developer. Struktur ini melibatkan pembagian perangkat lunak menjadi berbagai komponen yang saling berinteraksi. Tujuan dari struktur ini adalah untuk memfasilitasi berbagai tahapan dalam pengembangan, yang mencakup hal-hal berikut:

- Development
- Deploying
- Operation
- Maintenance

Terdapat salah satu jenis arsitektur yang umum digunakan dalam pengembangan perangkat lunak, yakni Clean Architecture. Arsitektur ini mengadopsi pendekatan pemisahan perangkat lunak ke dalam beberapa lapisan, di mana masing-masing lapisan memiliki tanggung jawab yang berbeda. Pada setidaknya satu lapisan terdapat tanggung jawab untuk logika bisnis (business logic), sementara lapisan lainnya bertanggung jawab atas antarmuka pengguna dan sistem.

Berdasarkan penjelasan sebelumnya, prinsip-prinsip dasar dari clean architecture antara lain sebagai berikut:

- Tidak bergantung pada framework tertentu. Hal ini memungkinkan developer menggunakan framework hanya sebagai alat bantu untuk pengembangan.
- Dapat diuji. Logika bisnis dapat diuji secara terpisah tanpa harus melibatkan antarmuka pengguna, basis data, server, dan sebagainya.
- Tidak bergantung pada desain antarmuka pengguna tertentu. Ini berarti desain antarmuka pengguna dapat diubah dengan mudah tanpa memengaruhi bagian lain dari perangkat lunak (seperti basis data, server, dan lain-lain).
- Tidak bergantung pada basis data tertentu. Kapan pun diperlukan, developer dapat mengubah basis data dari relasional ke non-relasional, atau sebaliknya, tanpa harus memengaruhi logika bisnis karena logika bisnis tersebut tidak bergantung pada basis data.

Adapun daftar lapisan atau layer pada clean architecture antara lain sebagai berikut:

- Entities (enterprise business rules)
- Use Cases (application business rules)
- Interface Adapters
- Frameworks & Drivers

## Entities (Enterprise Business Rules)

Lapisan Entities biasanya terdiri dari serangkaian struktur dan fungsi yang berhubungan dengan pengelolaan data, misalnya contoh operasi create, read, update, dan delete. Umumnya, lapisan ini juga disebut sebagai repository. Lapisan ini memiliki kemungkinan terkecil untuk mengalami perubahan eksternal dalam perangkat lunak dan tidak memiliki ketergantungan pada lapisan lain.

## Use Cases (Application Business Rules)

Lapisan Use Cases umumnya terdiri dari sekelompok fungsi yang merangkai logika bisnis pada perangkat lunak. Lapisan ini sering juga disebut sebagai service atau use case. Di dalam lapisan ini, alur data diimplementasikan dari dan menuju entities untuk mencapai tujuan dalam setiap aspek logika bisnis pada perangkat lunak. Lapisan ini mengimplementasikan fungsi-fungsi yang disediakan oleh lapisan Entities.

## Interface Adapters

Lapisan Interface Adapters biasanya terdiri dari kumpulan fungsi yang bertugas mengkonversi data dari format eksternal, seperti permintaan (request), ke dalam bentuk data yang dapat dikelola melalui lapisan Use Cases. Lapisan ini sering juga disebut sebagai controller atau handler. Lapisan ini bertugas mengimplementasikan fungsi-fungsi yang telah didefinisikan dalam lapisan Use Cases.

## Frameworks & Drivers

Lapisan Frameworks and Drivers umumnya terdiri dari kumpulan kode yang memberikan detail komponen atau teknologi yang diperlukan dalam perangkat lunak yang telah dikembangkan, seperti framework, pustaka (library), dan basis data.

### Kesimpulan

Clean Architecture adalah jenis arsitektur perangkat lunak yang membagi aplikasi menjadi beberapa lapisan berdasarkan tanggung jawabnya. Setiap lapisan memiliki perannya masing-masing: Entities untuk pengelolaan data, Use Cases untuk logika bisnis, Interface Adapters untuk konversi dari bentuk eksternal ke data yang dapat dikelola, dan Frameworks & Drivers untuk mengatur detail teknis seperti framework, pustaka, dan basis data yang digunakan.

Implementasi arsitektur ini memungkinkan pengembang untuk membangun perangkat lunak yang tidak tergantung pada antarmuka pengguna, basis data, atau kerangka kerja tertentu. Hal ini menghasilkan beberapa keunggulan, termasuk fleksibilitas, struktur yang terorganisir, kebersihan kode, dan kemudahan pemeliharaan. Dengan menjaga pemisahan yang jelas antara berbagai lapisan, Clean Architecture membantu memastikan skalabilitas dan keserbagunaan aplikasi di tengah perubahan yang mungkin terjadi di masa depan.

source : [Clean Architecture untuk Development REST-API ](https://blog.dot.co.id/clean-architecture-untuk-development-rest-api-yang-maintainable-9212e05fe38b)
