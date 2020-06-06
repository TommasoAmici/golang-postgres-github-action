-- initialize postgres database
-- drop all before initializing
drop table users;
-- create tables
create table if not exists users (id serial primary key, username varchar(128));
insert into users (username)
values ('User name');