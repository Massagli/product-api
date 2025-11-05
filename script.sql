create table product(
	ID serial primary key,
	product_name varchar(50) not null,
	price numeric(10, 2) not null
);
select * from product;
insert into product (product_name, price) values ('Pão', 1.50);

create table users(
	ID serial primary key,
	user_name varchar(100) not null,
	user_email varchar(100) not null
);
drop table users;

select * from users;
insert into users (user_name, user_email) values ('Adriel', 'adriel@gmail.com');