# System Design

## Diagram

Diagram adalah suatu representasi simbolis informasi dalam bentuk geometri 2 dimensi sesuai teknik visualisasi.

Berikut beberapa tool yang dapat digunakan dalam visualisasi Diagram

- smartdraw
- lucidchart
- whimsical
- draw.io
- visio

Beberapa software design yang sering digunakan

- Flowchart
- Use Case
- ERD (Entity Relationship Diagram)
- HLA (High Level Architecture)

## Horizontal Scaling and Vertical Scaling

### System Design Basics

Kapanpun kita mendesain sebuah sistem yang besar, kita perlu mempertimbangkan beberapa hal yaitu:

- Apa perbedaan beberapa arsitektur yang bisa kita gunakan
- Bagaimana setiap arsitektur itu dapat bekerja satu sama lain
- Bagaimana kita bisa menggunakannya dengan tepat

Mengerti beberapa konsep diatas dapat memberikan kita banyak keuntungan dalam memahami konsep sistem distribusi.

### Key Characteristics of Distrubted Systems

- Scability (kapabilitas sebuah sistem, dalam menangani kenaikan permintaan)
- Reliability (kapabilitas sebuah sistem, dalam memberikan layanan)
- Availability (kapabilitas sebuah sistem, operasional meskipun dalam kondisi sedang tidak normal)
- Efficiency (kapabilitas sebuah sistem, dalam hal efisiensi)
- Serviceability or Manageability (kapabilitas sebuah sistem, dalam hal kecepatan seberapa cepat sistem tersebut dapat diperbaiki dan dimaintenance)

## Job/Work Queue

Dalam design system, Job/Work Queue adalah antrian pekerjaan, sebuah struktur data yang dikelola oleh perangkat lunak penjadwal pekerjaan yang berisi pekerjaan yang harus dijalankan.

## Load Balancing

Load balancing merupakan proses pendistribusian traffic atau lalu lintas jaringan secara efisien ke dalam sekelompok server, atau yang lebih dikenal dengan server pool atau server farm. Load balancing ini berguna agar salah satu server dari website yang mendapatkan banyak lalu linta kunjungan tidak mengalami kelebihan beban.

## Monolithic and Microservices

Monolithic merupakan sebuah pendekatan tradisional dalam pembangunan sebuah aplikasi. Aplikasi monolitik terbentuk sebagai satu kesatuan kode yang tidak dapat dipisahkan satu dan lainnya. Salah satu karakter sistem arsitektur monolitik adalah saat pemrogram ingin melakukan perubahan pada sistem monolitik, pemrogram harus mengubah satu kesatuan kode secara menyeluruh dan bersamaan.

Microservices merupakan kebalikan dari monolithic, jika monolithic adalah sebuah arsitektur aplikasi secara kesatuan atau tunggal, maka Microservices terbagi menjadi unit pecahan yang lebih kecil dan spesifik. Setiap unitnya terpisah dan memiliki sistem beserta database sendiri untuk beroperasi dan menggunakan mekanisme API untuk terhubung dengan unit lainnya.

## SQL and NoSQL

SQL seperti yang kita semua sudah ketahui pada materi sebelumnya ialah bahasa yang digunakan untuk mengatur/mengelola data dalam database relasional. Database relasional menggunakan ‘relasi’ (yang biasanya disebut tabel) untuk menyimpan data dan mencocokkan data tersebut dengan memakai karakteristik umum di setiap dataset. Beberapa contoh database management system yang menggunakan SQL antara lain Oracle, Sybase, Microsoft SQL Server, PostgreSQL.

NoSQL merupakan databas yang tidak membutuhkan skema dan tidak memiliki relasi untuk setiap table. Semua bentuk dokumen dari NoSQL adalah JSON yang mudah dibaca dan dimengerti. NoSQL banyak diminati karena memilik peforma yang tinggi dan bersifat non-relasional sehingga dapat memakai berbagai model data. Beberapa contoh dari database NoSQL yaitu MongoDB, MarkLogic, CouchBase, CloudDB, dan Dynamo DB.

## Caching

Dalam komputasi, cache adalah lapisan penyimpanan data berkecepatan tinggi yang menyimpan sebagian data, biasanya bersifat sementara, sehingga permintaan di masa mendatang untuk data tersebut dilayani lebih cepat daripada yang mungkin dilakukan dengan mengakses lokasi penyimpanan utama data. Cache memungkinkan kita dapat menggunakan kembali data yang diambil atau dihitung sebelumnya secara efisien.

Bagaimana cara kerja caching ?
Data dalam cache umumnya disimpan dalam hardware akses cepat seperti RAM (Random Access Memory) dan juga dapat digunakan dalam hubungannya dengan komponen perangkat lunak. Tujuan utama cache adalah untuk meningkatkan kinerja pengambilan data dengan mengurangi kebutuhan untuk mengakses lapisan penyimpanan yang lebih lambat.
Menukar kapasitas dengan kecepatan, cache biasanya menyimpan sebagian data secara sementara, berbeda dengan database yang datanya biasanya lengkap dan tahan lama.

## Database Replication

Database Replication adalah sistem yang digunakan untuk menyalin, mendistribusikan data, serta melakukan sinkronisasi data antar database. Database Replication juga dapat membantu meningkatkan kinerja database dengan mempercepat waktu akses dan meminimalkan waktu downtime yang mungkin terjadi jika hanya ada satu database tunggal.

Pada dasarnya sistem replication membutuhkan minimal dua buah server untuk digunakan sebagai master dan slave. Dengan menggunakan teknik replikasi, data dapat didistribusikan ke server yang berbeda melalui koneksi jaringan lokal maupun internet.

Sebagai contoh, sebuah aplikasi biasanya dapat mengakses database, aplikasi akan dapat terus berfungsi jika server lokal mengalami kegagalan/down, karena terdapat server database replikasi lain yang tetap berjalan dan berisi data yang sama sebagai backup apabila server utama mengalami down.

## Database Indexing

Pada database, index merupakan sebuah struktur data yang berisi kumpulan keys beserta referensinya ke actual data di table. Tujuannya untuk mempercepat proses penentuan lokasi data tanpa melakukan pencarian secara penuh ke seluruh data (full scan). Fitur index sangat membantu dalam mengoptimalkan proses query, terutama dengan data yang besar. Satu-satunya additional cost pada penggunaan fitur index adalah bertambahnya ukuran data secara total, karena index memiliki ukuran data tersendiri, meskipun relatif jauh lebih kecil daripada data asli.
