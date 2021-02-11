![Github CI/CD](https://img.shields.io/github/workflow/status/denchick/trademarks/Go)
![Go Report](https://goreportcard.com/badge/github.com/denchick/trademarks)
![Repository Top Language](https://img.shields.io/github/languages/top/denchick/trademarks)
![Scrutinizer Code Quality](https://img.shields.io/scrutinizer/quality/g/denchick/trademarks/master)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/denchick/trademarks)
![Codacy Grade](https://img.shields.io/codacy/grade/c9467ed47e064b1981e53862d0286d65)
![Github Repository Size](https://img.shields.io/github/repo-size/denchick/trademarks)
![Github Open Issues](https://img.shields.io/github/issues/denchick/trademarks)
![Lines of code](https://img.shields.io/tokei/lines/github/denchick/trademarks)
![License](https://img.shields.io/badge/license-MIT-green)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/denchick/trademarks)
![GitHub last commit](https://img.shields.io/github/last-commit/denchick/trademarks)
![GitHub contributors](https://img.shields.io/github/contributors/denchick/trademarks)
![Simply the best ;)](https://img.shields.io/badge/simply-the%20best%20%3B%29-orange)

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
