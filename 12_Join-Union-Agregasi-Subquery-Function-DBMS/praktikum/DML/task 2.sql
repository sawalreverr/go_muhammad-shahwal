-- Task 2

-- task 2 - A
select * from users where gender = 'M'; 

-- task 2 - B
select * from products where id = 3;

-- task 2 - C
select * from users where created_at > current_date - interval 7 day and name like '%a%';

-- task 2 - D
select count(*) from users where gender = 'F';

-- task 2 - E
select * from users order by name asc;

-- task 2 - F
select * from products limit 5;