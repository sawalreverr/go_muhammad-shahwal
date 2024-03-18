-- create database task_12;

use task_12;

create table users (
	id int(11) primary key auto_increment,
  name varchar(255),
  status smallint,
  dob date,
  gender char(1),
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp on update current_timestamp 
);

create table payment_methods (
	id int(11) primary key auto_increment,
  name varchar(255),
  status smallint,
	created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp on update current_timestamp
);

create table transactions (
	id int(11) primary key auto_increment,
  user_id int(11),
  payment_method_id int(11),
  status varchar(10),
  total_qty int(11),
  total_price numeric(25,2),
	created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp on update current_timestamp,
  constraint fk_transcation_user foreign key (user_id) references users(id),
  constraint fk_transcation_payment_method foreign key (payment_method_id) references payment_methods(id)
);

create table product_types (
	id int(11) primary key auto_increment,
  name varchar(255),
	created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp on update current_timestamp
);

create table operators (
	id int(11) primary key auto_increment,
  name varchar(255),
	created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp on update current_timestamp
);

create table products (
	id int(11) primary key auto_increment,
  product_type_id int(11),
  operator_id int(11),
  code varchar(50),
  name varchar(100),
  status smallint,
	created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp on update current_timestamp,
  constraint fk_product_type foreign key (product_type_id) references product_types(id),
  constraint fk_product_operator foreign key (operator_id) references operators(id)
);

create table product_descriptions (
	id int(11) primary key auto_increment,
  description text,
	created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp on update current_timestamp
);

create table transaction_details (
	transcation_id int(11),
  product_id int(11),
  status varchar(10),
  qty int(11),
  price numeric(25,2),
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp on update current_timestamp,
  primary key (transcation_id, product_id),
  constraint fk_transaction_details_transaction foreign key (transcation_id) references transactions(id),
  constraint fk_transaction_details_product foreign key (product_id) references products(id)
);