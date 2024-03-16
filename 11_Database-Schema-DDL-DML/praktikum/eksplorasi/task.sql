# create database knowledge_management;

use knowledge_management;

create table users (
	id integer primary key auto_increment,
  username varchar(255),
  password varchar(255),
  email varchar(255)
);

create table notes (
	id integer primary key auto_increment,
  title varchar(255),
  content text,
  category varchar(255),
  user_id integer,
  constraint fk_note_user foreign key (user_id) references users(id)
);

create table tasks (
	id integer primary key auto_increment,
  title varchar(255),
  description text,
  status varchar(255),
  deleted_at datetime,
  user_id integer,
  constraint fk_task_user foreign key (user_id) references users(id)
);

create table medias (
	id integer primary key auto_increment,
  type varchar(255),
  file_path varchar(255),
  note_id integer,
  constraint fk_media_note foreign key (note_id) references notes(id)
);

create table pages (
	id integer primary key auto_increment,
  title varchar(255),
  content text,
  user_id integer,
  constraint fk_page_user foreign key (user_id) references users(id)
);

create table page_medias (
	id integer primary key auto_increment,
  media_id integer,
  page_id integer,
  constraint fk_page_medias_media foreign key (media_id) references medias(id),
  constraint fk_page_medias_page foreign key (page_id) references pages(id)
);

desc users;
desc notes;
desc tasks;
desc medias;
desc pages;
desc page_medias;