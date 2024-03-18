-- Data Manipulation Language (DML)

-- Task 1

-- task 1 - A
insert into operators (name) 
values 
	('Salim Suhartono'), 
  ('Melinda Suhartono'), 
  ('Ilham Suhartono'), 
  ('Kurniawan Suhartono'), 
  ('Rizki Suhartono');

-- task 1 - B
insert into product_types (name)
values
	('Elektronik'),
  ('Fashion Pria'),
	('Fashion Wanita');

-- task 1 - C
insert into products (product_type_id, operator_id, code, name, status)
values
	(1, 3, 'E-1000', 'Air Cooler', 1),
  (1, 3, 'E-1001', 'Mesin Cuci', 1);
  
-- task 1 - D
insert into products (product_type_id, operator_id, code, name, status)
values
	(2, 1, 'FP-1000', 'Kemeja', 1),
  (2, 1, 'FP-1001', 'Sepatu', 1),
  (2, 1, 'FP-1002', 'Topi', 1);

-- task 1 - E
insert into products (product_type_id, operator_id, code, name, status)
values
	(3, 4, 'FW-1000', 'Rok', 1),
  (3, 4, 'FW-1001', 'Legging', 1),
  (3, 4, 'FW-1002', 'Celana Panjang', 1);

-- task 1 - F
insert into product_descriptions (description)
values
	('Air Coooler adalah metode pembuangan panas'),
  ('Mesin cuci adalah sebuah mesin yang dirancang untuk membersihkan pakaian dan tekstil rumah tangga lainnya seperti handuk dan sprai.'),
  ('Kemeja atau kamisa (dari bahasa Portugis: camisa) atau kelambi (dari bahasa Jawa) adalah salah satu jenis baju atau atasan yang terutama digunakan untuk pria.'),
  ('Sepatu merupakan alas kaki yang bersifat universal.'),
  ('Topi adalah suatu jenis penutup kepala.'),
  ('Rok adalah busana wanita yang dipakai pada badan bagian bawah, mulai dari pinggang dengan panjang bervariasi sesuai model yang umumnya Page 11 dibuat dengan cara dijahit bagian sisi.'),
  ('Legging adalah beberapa jenis pakaian kaki yang bervariasi dari tahun ke tahun. Penggunaan modern dari tahun 1960-an dan seterusnya telah merujuk pada pakaian ketat elastis bertingkat tinggi yang dikenakan di atas kaki biasanya oleh wanita, seperti penghangat kaki atau celana ketat.'),
  ('Celana panjang adalah pakaian luar untuk menutupi badan dari pinggang sampai mata kaki, yang mana pada bagian kaki dipisahkan antara kaki kiri dan kaki kanan');

-- task 1 - G
insert into payment_methods (name, status)
values
	('Credit Card', 1),
  ('BCA Virtual Account', 1),
  ('Gopay', 1);

-- task 1 - H
insert into users (name, status, dob, gender)
values
	('Ilham Kurniawan', 1, '2001-01-01', 'M'),
  ('Budi Kurniawan', 1, '2002-02-02', 'F'),
  ('Acep Kurniawan', 1, '2003-03-03', 'M'),
  ('Jafar Kurniawan', 1, '2004-04-04', 'F'),
  ('Rudi Kurniawan', 1, '2005-05-05', 'M');
   
-- task 1 - I
insert into transactions (user_id, payment_method_id, status, total_qty, total_price)
values
	(1, 2, 'payed', 3, 200000.00),
	(2, 1, 'not payed', 5, 500000.00),
  (3, 3, 'payed', 1, 150000.00);

-- task 1 - J
insert into transaction_details (transcation_id, product_id, status, qty, price)
values
	(1, 6, 'payed', '3', '200000.00'),
  (2, 5, 'not payed', '5', '500000.00'),
  (3, 8, 'payed', '1', '150000.00');


  