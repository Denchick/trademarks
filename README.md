# trademarks

// TODO description

## How it works

Idea is download all XMLs, parse it, convert to CSV and load into database.

// TODO

## How to run

// TODO how to download trademarks registered in 2019
0. Download archives from ftp://ftp.euipo.europa.eu/ and unzip it to somewhere
1. Start services with `docker-compose up -d `
2. Call `make` and `./xml2csv --directory /path/to/xmls`. Now you have `trademarks.csv`.
3. Time to load trademarks into database: `cat trademarks.csv | psql --host 0.0.0.0 --port 5432 --user postgres --dbname db_trademarks -c "COPY db_trademarks FROM STDIN WITH DELIMITER as ';' CSV HEADER"`

## Server

- Exact trademark could be search like this: `http://localhost:1323/v1/trademarks?name=vacuumlabs`
- 'nearest' trademarks  could be search like this: `http://localhost:1323/v1/trademarks?name=vacuumlabs&fuzzily=true`
