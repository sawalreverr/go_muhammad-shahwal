use lowongan_pekerjaan;

create table users (
	id integer primary key auto_increment,
  name varchar(255),
  email varchar(255),
  address varchar(255),
  description text
);

create table work_experiences (
	id integer primary key auto_increment,
  user_id integer,
  company varchar(255),
  role varchar(255),
  start_date datetime,
  end_date datetime,
  description text,
  constraint fk_work_experience_user foreign key (user_id) references users(id)
);

create table recruiters (
	id integer primary key auto_increment,
  name varchar(255),
  email varchar(255),
  address varchar(255),
  description text
);

create table jobs (
	id integer primary key auto_increment,
  recruiter_id integer,
  title varchar(255),
  description text,
 	category varchar(255),
  status varchar(255),
  constraint fk_job_recruiter foreign key (recruiter_id) references recruiters(id) 
);

create table applications (
	id integer primary key auto_increment,
  user_id integer,
  job_id integer,
  cv_file varchar(255),
  cover_letter text,
  status varchar(255),
  constraint fk_user_application foreign key (user_id) references users(id),
  constraint fk_job_application foreign key (job_id) references jobs(id)
);

desc users;
desc work_experiences;
desc recruiters;
desc jobs;
desc applications;