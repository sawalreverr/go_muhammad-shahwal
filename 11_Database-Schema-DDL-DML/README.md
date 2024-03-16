# Database: Schema & Data Definition Language

## Introduction Database

Database atau basis data adalah kumpulan data yang dikelola sedemikian rupa berdasarkan ketentuan tertentu yang saling berhubungan sehingga mudah dalam pengelolaannya. Melalui pengelolaan tersebut pengguna dapat memperoleh kemudahan dalam mencari informasi, menyimpan informasi dan membuang informasi.

## Jenis-jenis Database

### Operational Database

Operational Database atau biasa disebut dengan database OLTP (On Line Transaction Processing), berguna untuk mengelola data yang dinamis secara langsung atau real-time. Jenis ini memungkinkan para pengguna dapat melakukan, melihat, dan memodifikasi data. Modifikasi tersebut bisa berupa mengubah, menambah, menghapus data secara langsung melalui perangkat keras yang digunakan.

- JSON (JavaScript Object Notation)
- XML (Extensible Markup Language)

### Database Warehouse

Database Warehouse adalah sistem basis data yang biasa digunakan untuk pelaporan dan analisis data. Sistem ini dianggap sebagai komponen inti dari business intelligence. Database Warehouse merupakan repositori sentral data yang terpadu dari satu atau lebih sumber yang berbeda. Database tersebut juga menyimpan data terkini dan historis dengan satu tempat yang digunakan untuk membuat laporan analisis.

- Microsoft SQL Server

### Distributed Database

Distributed Database adalah basis data yang perangkat penyimpanannya tidak terpasang pada perangkat komputer yang sama. Basis data tersebut disimpan di beberapa perangkat komputer yang terletak di tempat yang sama atau tersebar melalui jaringan komputer lainnya yang saling berhubungan.

- Microsoft Access (Office)

### Relational Database

Relational Database atau basis data relasional adalah basis data yang mengorganisir berdasarkan model hubungan data. Banyak sekali perangkat lunak yang menggunakan sistem ini untuk mengatur dan memelihara basis data melalui hubungan setiap data. Umumnya, semua sistem menggunakan Structured Query Language (SQL) sebagai bahasa pemrograman untuk pemeliharaan basis data dan query.

- MySQL
- PostgreSQL
- MariaDB
- MongoDB
- Oracle Database
- SAP HANA

### End-User Database

SQLite adalah sistem manajemen basis data yang ada pada library pemrograman C. Berbeda dengan sistem lainnya, SQLite bukan merupakan mesin database client server. SQLite tertanam ke dalam program akhir sehingga cocok digunakan dalam mendukung penyimpanan data akhir end user.

## Database Relationship

### ONE TO ONE RELATIONSHIP

One-To-One Relationship, adalah hanya satu dari masing- masing entity yang saling berhubungan atau berelasi. (1 user hanya memiliki 1 foto profil).

### ONE TO MANY RELATIONSHIP

One-To-Many Relationship, adalah dimana satu attribute dari satu entity bisa berhubungan dengan dua atau lebih attribute dari entity yang lain. Contoh nya adalah mata pelajaran bisa diajar oleh beberapa guru, sedangkan satu guru hanya mengajar pada satu mata pelajaran saja.

### MANY TO MANY RELATIONSHIP

Many–To-many Relationship, adalah suatu relasi dimana satu attribute dari satu entity dapat memiliki hubungan dengan dua atau lebih attribute dari entity yang lain, begitupun dengan sebaliknya. Contohnya satu mata pelajaran dapat diikuti oleh banyak siswa, dan begitupun sebaliknya satu siswa mendapat mata pelajaran bisa lebih dari satu mata pelajaran.

## Jenis-jenis Perintah SQL

### DDL (Data Definition Language)

```sql
# DDL Statement

CREATE DATABASE database_name;
USE database_name;
CREATE TABLE....
DROP TABLE....
RENAME TABLE....

# Create Table With Is Schema
CREATE TABLE table_name(
    column1 data_type PRIMARY KEY,
    column1 data_type FOREIGN KEY,
    ....
    column data_type,
    PRIMARY KEY(one or more columns)
);

# Modify Table Schema
ALTER TABLE table_name
ADD COLUMN column_name data_type;
```

#### Tipe data MySQL

- Num
- Huruf
- Date

#### DML (Data Manipulation Language)

Perintah yang digunakan untuk memanipulasi data dalam tabel dari suatu database.

Statement Operation:

- INSERT
- SELECT
- UPDATE
- DELETE

```sql
# INSERT
INSERT INTO USERS (username, fullname, status, gender, email, password, location)
VALUES ('hadildo', 'Nurhadi Aldo', 1, 'm', 'hadildo@tronjaltronjol.com', '123456', 'Jawa Timur');

# SELECT
SELECT * FROM USERS;
SELECT username, fullname FROM USERS WHERE id=1;
SELECT username, fullname FROM USERS WHERE fullname IS NOT NULL;

# UPDATE
UPDATE users SET fullname='Nurhadi Aldo Tronjal Tronjol' WHERE id=1;

# DELETE
DELETE FROM users WHERE id=1;
```

DML Statement

- LIKE / BETWEEN
- AND / OR
- ORDER BY
- LIMIT

```sql
# LIKE / BETWEEN
SELECT user_id, type, message, parent_id FROM tweets WHERE message LIKE 'H%';
SELECT user_id, type, message, parent_id FROM tweets WHERE user_id BETWEEN 1 AND 3;

# AND / OR
SELECT user_id, type, message, parent_id FROM tweets WHERE message LIKE 'H%' OR user_id BETWEEN 1 AND 3;
SELECT user_id, type, message, parent_id FROM tweets WHERE message LIKE 'H%' AND user_id BETWEEN 1 AND 2;

# ORDER BY
SELECT id, user_id, type, message, parent_id FROM tweets WHERE message LIKE 'H%' OR user_id BETWEEN 1 AND 2 ORDER BY id DESC;

# LIMIT
SELECT user_id, type, message, parent_id FROM tweets WHERE message LIKE 'H%' AND user_id BETWEEN 1 AND 3 ORDER BY 1 DESC LIMIT 2;
```

JOIN (sebuah klausa untuk mengkombinasikan record dari dua atau lebih tabel)

- INNER
- LEFT
- RIGHT

```sql
# INNER JOIN
SELECT t.message FROM users u INNER JOIN tweets t ON u.id = t.user_id;

# LEFT JOIN
SELECT u.username, t.message FROM users u LEFT JOIN tweets t ON u.id = t.user_id;

# RIGHT JOIN
SELECT u.username, t.message FROM users u RIGHT JOIN tweets t ON u.id = t.user_id;
```

UNION

```sql
SELECT username, fullname FROM users WHERE id=1 UNION SELECT username, fullname FROM users WHERE id=2;
```

AGGREGATE \
Fungsi Agregasi adalah fungsi dimana nilai beberapa baris dikelompokkan bersama untuk membentuk nilai ringkasan tunggal.

- MIN
- MAX
- SUM
- AVG
- COUNT
- HAVING

```sql
# MIN
SELECT MIN(id) AS id FROM users;
SELECT MIN(created_at) AS created_at FROM users;

# MAX
SELECT MAX(id) FROM users
SELECT MAX(created_at) FROM users;

# SUM
SELECT SUM(favourite_count) FROM tweets WHERE user_id=1;

# AVG
SELECT AVG(favourite_count) FROM tweets WHERE user_id=1;

# COUNT
SELECT COUNT(1) FROM tweets WHERE user_id=1;

# HAVING
SELECT user_id FROM tweets GROUP BY user_id HAVING SUM(favourite_count) > 2;
```

FUNCTION \

```sql
DELIMITER $$
CREATE FUNCTION sf_count_tweet_peruser (user_id_p int) RETURNS INT DETERMINISTIC
BEGIN
DECLARE total INT;
SELECT COUNT(*) INTO total FROM tweets WHERE user_id = user_id_p AND type='tweets';
RETURN total;
END$$
DELIMITER;
```

TRIGGER FUNCTION

```sql
DELIMITER $$
CREATE TRIGGER delete_all_data_user
BEFORE DELETE ON users FOR EACH ROW
BEGIN
-- declare variables
DECLARE v_user_id INT;
SET v_user_id=OLD.id;
-- trigger code
DELETE FROM tweets WHERE user_id = v_user_id;
DELETE FROM user_followers WHERE user_id = v_user_id;
END$$
```
