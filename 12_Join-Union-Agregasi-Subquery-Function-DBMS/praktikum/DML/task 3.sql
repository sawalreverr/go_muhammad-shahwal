-- Task 3

-- task 3 - A
update products set name='product dummy' where id = 1; 
select * from products;

-- task 3 - B
update transaction_details set qty=3 where product_id = 1;
select * from transaction_details;
