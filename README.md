# trademarks

In the developed world, businesses can protect their brands by trademark registration. During such a process, the user needs to check whether their trademark is not already taken by someone else—this is roughly similar to how domain registration works. Your job is to write a simple backend-only validation service operating on top of a given trademark database.

## How it works

The idea was to download all XMLs, parse it, convert them to CSV, and load them into PostgreSQL. The back-end implements a simple REST API on Go.

## Requirements

- [go 1.15](https://golang.org/doc/install)
- docker

## Before start

Before the start, we should download and prepare data for the database. If I right understand, we need files inside the `ftp://ftp.euipo.europa.eu/Trademark/Full/2019` folder starting with `EUTMS_`: 

- EUTMS_20201118_0001.zip, 
- ..., 
- EUTMS_20201118_0006.zip.

Let's download and unzip them:

```bash
TRADEMARKS_DIR=$(mktemp -d)
wget -m --ftp-user=opendata --ftp-password=kagar1n --directory-prefix $TRADEMARKS_DIR --no-directories ftp://ftp.euipo.europa.eu/Trademark/Full/2019/EUTMS_20201118_000{1..6}.zip
cd $TRADEMARKS_DIR
unzip "*.zip"
rm *.zip
```

## How to run

1. Start services with `docker-compose up -d `
2. Call `make` and `./xml2csv --directory $TRADEMARKS_DIR`. Now you have `trademarks.csv`.
3. Time to load trademarks into database: `cat trademarks.csv | psql --host 0.0.0.0 --port 5432 --user postgres --dbname db_trademarks -c "COPY db_trademarks FROM STDIN WITH DELIMITER as ';' CSV HEADER"`

You are ready to search!

- To search exact trademark: `http://localhost:1323/v1/trademarks?name=vacuumlabs`
- To search nearest trademark: `http://localhost:1323/v1/trademarks?name=vacuumlabs&fuzzily=true`