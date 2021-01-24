CREATE USER trademarks WITH PASSWORD 'trademarks';
CREATE DATABASE trademarksdb;
GRANT ALL PRIVILEGES ON DATABASE trademarksdb TO trademarks;
CREATE USER volkov WITH PASSWORD 'trademarks';
GRANT ALL PRIVILEGES ON DATABASE trademarksdb TO volkov;
