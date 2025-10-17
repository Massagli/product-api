create table product(
	ID serial primary key,
	product_name varchar(50) not null,
	price numeric(10, 2) not null
);

select * from product;

insert into product (product_name, price) values ('Pão', 1.50);