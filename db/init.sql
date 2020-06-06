-- initialize postgres database
DROP DATABASE IF EXISTS test_golang;
DROP USER IF EXISTS test_golang;
CREATE USER test_golang CREATEDB PASSWORD 'test_golang';
CREATE DATABASE test_golang OWNER test_golang;
-- drop all before initializing
DROP TABLE users;
-- create tables
CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, username VARCHAR(128));
INSERT INTO users (username)
VALUES ('User name');