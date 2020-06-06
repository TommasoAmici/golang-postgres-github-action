-- initialize postgres database
-- drop all before initializing
DROP TABLE users;
-- create tables
CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, username VARCHAR(128));
INSERT INTO users (username)
VALUES ('User name');