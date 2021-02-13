![Github CI/CD](https://img.shields.io/github/workflow/status/denchick/trademarks/Go?style=for-the-badge&logo=github)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/denchick/trademarks?style=for-the-badge)
![Go Report](https://goreportcard.com/badge/github.com/denchick/trademarks?style=for-the-badge)
![Github Repository Size](https://img.shields.io/github/repo-size/denchick/trademarks?style=for-the-badge)
![Lines of code](https://img.shields.io/tokei/lines/github/denchick/trademarks?style=for-the-badge)

# trademarks ™️

<img align="right" src="./assets/gopher.png">

Businesses can protect their brands by trademark registration. During such a process, the user needs to check whether their trademark is not already taken by someone else—this is roughly similar to how domain registration works. There is a special organization responsible for managing the EU trade mark and the registered Community design - *European Union Intellectual Property Office*. They put all the information about the trademark registration process in the public domain, so we can use them.

I wrote the service that check if there was already registered exact or similar trademark.

## Requirements

- [go 1.15](https://golang.org/doc/install)
- docker

## Before the start

> You can skip that step and use `./assets/trademarks.csv`

Before the start, we should download and prepare data for the database. We need files inside the `ftp://ftp.euipo.europa.eu/Trademark/Full/2019`(username: `opendata` password: `kagar1n`) folder starting with `EUTMS_`.

Let's download and unzip them:

```bash
TRADEMARKS_DIR=$(mktemp -d)
wget -m 
    --ftp-user=opendata 
    --ftp-password=kagar1n 
    --directory-prefix $TRADEMARKS_DIR 
    --no-directories
    ftp://ftp.euipo.europa.eu/Trademark/Full/2019/EUTMS_20201118_000{1..6}.zip
cd $TRADEMARKS_DIR
unzip "*.zip"
rm *.zip
```

## How to run

1. Start services with `docker-compose up -d `
2. Call `make` and `./xml2csv --directory $TRADEMARKS_DIR`. Now you have `trademarks.csv`.
3. Time to load trademarks into database: `cat trademarks.csv | psql --host 0.0.0.0 --port 5432 --user postgres --dbname db_trademarks -c "COPY db_trademarks FROM STDIN WITH DELIMITER as ';' CSV HEADER"`

Database is ready!

- To search exact trademark: `http://localhost:1323/v1/trademarks?name=UNISON`
- To search nearest trademark: `http://localhost:1323/v1/trademarks?name=uniso&similar=true`
