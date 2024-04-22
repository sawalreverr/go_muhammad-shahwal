# Docker

## What is Docker ?

Docker adalah layanan yang menyediakan kemampuan untuk mengemas dan menjalankan sebuah aplikasi dalam sebuah lingkungan terisolasi yang disebut dengan container. Dengan adanya isolasi dan keamanan yang memadai memungkinkan kamu untuk menjalankan banyak container di waktu yang bersamaan pada host tertentu.

Docker ini diperkenalkan pada tahun 2013 oleh Solomon Hykes pada acara PyCon. Beberapa bulan setelahnya docker secara resmi diluncurkan, tepatnya pada tahun 2014. Semenjak itu docker menjadi sangat populer di kalangan developer luar negeri, tetapi belum terlalu populer di Indonesia.

## Fungsi Docker

Fungsi utama penggunaan docker untuk pengembangan aplikasi adalah:

- Memudahkan Pengembangan

  Docker berfungsi untuk memudahkan pekerjaan pengembang saat mengembangkan aplikasi. Hal ini bisa terjadi karena Docker lebih hemat resource dam bisa menyediakan environment lebih stabil dijalankan pada semua perangkat seperti komputer pribadi dan cloud server.

- Memudahkan Pengembangan Kode Pipeline

  Docker ini berfungsi untuk memudahkan untuk mengatasi solusi permasalahan pengelolaan kode pipeline. Pengembang dapat memanfaatkan docker untuk melakukan pengujian pipeline code.

- Menyederhanakan Konfigurasi

  Docker mempunyai konfigurasi yang sederhana jadi bisa disesuaikan dengan aplikasi yang Anda jalankan. Bedanya fungsi docker dengan VM adalah docker ini tidak mempunyai overhead jadi bisa menjalankan aplikasi yang diuji tidak perlu konfigurasi tambahan.

- Mendukung Multitenancy

  Fungsi dari docker adalah membuat para pengembang terbantu dengan kemampuan docker yang dapat mendukung pembuatan aplikasi dengan struktur multi-tenance. Contohnya saat perusahaan Anda akan mengembangkan program tersebut untuk mengaplikasikan IoT.

  Adanya docker, Anda dapat membuat environment lebih dari satu tanpa penulisan ulang di aplikasi utama dan menjalankan objek aplikasi di setiap tenant. Apabila dilakukan dengan manual bisa membuang banyak waktu sehingga Anda harus memanfaatkan docker.

- Digunakan untuk Debugging

  Pengembang bisa memanfaatkan docker untuk debugging secara cepat. Proses ini pada Sandbox akan memerlukan waktu 1 menit saja jadi tergolong sangat cepat.

- Meningkatkan Sumber Daya secara Cepat

  Dengan bantuan docker, pengembang bisa menjalankan sejumlah layanan sekaligus jadi bisa menunjang produktivitas saat melakukan pengembangan aplikasi dengan skala cukup besar.

## Fitur-fitur docker

- Docker container

  Docker container merupakan tempat untuk menjalankan dan mengemas aplikasi. Tempat ini termasuk runtime, kode, tools dan pengaturan. Container hanya dapat mengakses source yang sudah ditentukan pada docker image.

- Docker image

  Docker image merupakan kumpulan file untuk menunjang sebuah aplikasi. Docker image termasuk komponen untuk menginstruksikan docker server tentang syarat cara container docker dibuat.

- Docker daemon

  Docker daemon merupakan tempat untuk mengelola docker image, network, container dan volume. Tempat ini menyediakan command-line-interface (CLI) sisi client supaya user bisa berinteraksi dengan daemon melalui request dari Docker API yang nantinya diproses oleh sistem.

- Docker Engine RestAPI

  Docker Engine RestAPI merupakan komponen yang dapat diakses client dari HTTP agar bisa berinteraksi dengan docker daemon.

- Docker registry

  Docket registry merupakan tempat menyimpan docker image yang nantinya memberikan output sesuai perintah yang diinginkan.

- Docker host

  Docker host merupakan komponen yang memiliki tugas untuk menerima perintah dari docker client serta menyediakan environment untuk menjalankan suatu aplikasi.

- Docker hub

  Docker hub merupakan layanan yang bisa dimanfaatkan untuk sejumlah container image.

## Cara Kerja Docker

Container pada docker adalah lingkungan virtual ringan dan portabel yang mampu berjalan di platform apa saja jadi akan memudahkan untuk pengembangan, pengujian dan deploy aplikasi. Gambaran singkat mengenai cara kerjanya adalah docker berjalan dengan sistem operasi host, seperti Mac, Windows atau Linux. Docker akan memakai arsitektur client-server dan klien berkomunikasi dengan docker daemon yang memiliki tanggung jawab membangun, mengelola dan menjalankan container.

Docker image adalah read-only template yang isinya seluruh file dan dependensi sesuai kebutuhan agar bisa menjalankan aplikasi. Gambar ini dibuat memakai Dockerfile, yaitu skrip yang dapat menentukan lingkungan serta konfigurasi aplikasi. Docker image disimpan pada registry publik seperti registry pribadi atau Docker hub.

Agar bisa menjalankan aplikasi pada container, Anda harus membuat Docker container terlebih dahulu dari gambar. Kontainer merupakan instansi tulis dari gambar dan bisa berjalan dengan terisolasi dari sistem host atau kontainer yang lain.

Docker menyediakan interface baris perintah atau CLI dan API untuk mengelola gambar, kontainer serta sumber daya yang lainnya. Docker juga telah mendukung alat orkestrasi kontainer seperti Kubernetes dan Docker Swarm yang bisa membantu dalam pengelolaan klaster kontainer di sejumlah host.

Secara menyeluruh, kemampuan docker dapat menyederhanakan proses deploy, membangun dan mengelola aplikasi dengan mengisolasi mereka ke kontainer serta menyediakan platform sesuai standar untuk bisa menjalankannya dengan baik.

### Docker Basic Commands

```bash
# pull images
docker pull redis:latest

# create image / build image
docker container create -it --name redis1 redis:latest

# check all container
docker container ls -a
docker ps -a

# start / stop container
docker container start redis1
docker container stop redis1

# check detail container
docker inspect redis1

# check logs container
docker logs redis1

# menjalankan command / exec pada container
docker exec -it redis1 redis-cli

# menghapus container
docker container rm redis1
docker container rm -f redis1 # hapus secara paksa

# list all images
docker images
```

### Docker Pada Go App

```bash
# build image
docker build -t go-mini-api:1.0.0 .

# run image
docker run -itd --name mini-api -p 1323:1323 go-mini-api:1.0.0

# stop image
docker stop mini-api

# tagging dan push tag ke docker hub
docker tag go-mini-api:1.0.0 sawalrever23/go-mini-api:1.0.0
docker push sawalrever23/go-mini-api:1.0.0

# pull dari docker hub
docker pull sawalrever23/go-mini-api

# hapus image
docker image rm sawalrever23/go-mini-api:1.0.0
```

### Docker Volume

Digunakan untuk menyimpan semua data yang pernah ada pada docker container, sebagai contoh database.

```bash
# run and exec
docker run -itd --name cobadb -p 3306:3306 -e MARIADB_ALLOW_EMPTY_ROOT_PASSWORD=yes mariadb:latest
docker exec -it cobadb mariadb

docker container stop cobadb
docker container start cobadb
docker exec -it cobadb mariadb

docker container rm -f cobadb

# penggunaan volume
mkdir cobavolume
docker run -itd --name cobadb -p 3306:3306 -e MARIADB_ALLOW_EMPTY_ROOT_PASSWORD=yes -v /home/nadir/cobavolume:/var/lib/mysql mariadb:latest
docker logs cobadb
```

### Docker Network

Digunakan ketika kita membutuhkan network atau host yang sama untuk container lainnya, seperti akses database ke api

```bash
# create network and list all network
docker network create apinetwork
docker network ls

# set network after run
docker network connect apinetwork dbku
docker network connect apinetwork coba-api

# set network before run
docker run -itd --name dbku -p 3306:3306 -e MARIADB_ALLOW_EMPTY_ROOT_PASSWORD=yes --network apinetwork mariadb:latest
docker run -itd --name coba-api -p 1323:1323 --network apinetwork go-mini-api:1.0.0
```

### Docker Compose

Digunakan untuk membuild beberapa container secara langsung.

```bash
# run and build
docker compose up -d

# down and removed
docker compose down
```
