# create database restaurant;

use restaurant;

create table restaurant_types (
	id integer primary key auto_increment,
  name varchar(255)
);

create table restaurants (
	id integer primary key auto_increment,
  name varchar(255),
  address text,
  restaurant_type_id integer,
  constraint fk_restaurant_type foreign key (restaurant_type_id) references restaurant_types(id)
);

create table menu_types (
	id integer primary key auto_increment,
  name varchar(255)
);

create table menus (
  id integer primary key auto_increment,
  name varchar(255),
  description text,
  price integer,
  menu_type_id integer,
  restaurant_id integer,
  constraint fk_menu_type foreign key (menu_type_id) references menu_types(id),
	constraint fk_restaurant foreign key (restaurant_id) references restaurants(id)
);

create table users (
	id integer primary key auto_increment,
  username varchar(255),
  email varchar(255),
  password varchar(255)
);

create table user_reviews (
	id integer primary key auto_increment,
  rating float,
  description text,
  restaurant_id integer,
  user_id integer,
  constraint fk_restaurant_review foreign key (restaurant_id) references restaurants(id),
  constraint fk_user_review foreign key (user_id) references users(id)
);

desc restaurants;
desc restaurant_types;

desc menus;
desc menu_types;

desc users;
desc user_reviews;
