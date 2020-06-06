#!/bin/bash
set -e

echo running init.sh

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    -- initialize postgres database
    -- drop all before initializing
    DROP TABLE users;
    -- create tables
    CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, username VARCHAR(128));
    INSERT INTO users (username)
    VALUES ('User name');
EOSQL