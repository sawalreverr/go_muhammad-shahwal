-- Join, Union, Sub query, Function

-- task 1
select t.* from transactions t inner join users u on t.user_id = u.id where u.id in (1,2);

-- task 2
select sum(total_price) from transactions where user_id = 1;

-- task 3
select count(*) from transactions t 
join transaction_details td on t.id = td.transcation_id
join products p on td.product_id = p.id
join product_types pt on p.product_type_id = pt.id
where pt.id = 2;

-- task 4
select p.*, pt.name as product_type_name from products p join product_types pt on p.product_type_id = pt.id;

-- task 5
select t.*, p.name as product_name, u.name as user_name from transactions t
join transaction_details td on t.id = td.transcation_id
join products p on td.product_id = p.id
join users u on t.user_id = u.id;

-- task 6
create function delete_td() returns trigger
begin
	delete from transaction_details where transcation_id = old.id;
	return old;
end;

-- trigger
create trigger delete_t 
after delete on transactions
for each row
begin 
	call delete_td();
end;

-- task 7
create function update_tqty() return trigger
begin
	declare qty_total int;
  select sum(qty) into qty_total from transaction_details where transcation_id = old.transaction_id;
  update transactions set total_qty = qty_total where id = old.transaction_id;
  return old;
end;

create trigger delete_td
after delete on transaction_details
for each row
begin 
	call update_tqty();
end;

-- task 8
select * from products where id not in (select distinct product_id from transaction_details);
