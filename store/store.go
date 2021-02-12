package store

import (
	"log"
	"time"

	"github.com/denchick/trademarks/config"
	"github.com/denchick/trademarks/store/repositories"
	"github.com/go-pg/pg/v10"
	"github.com/pkg/errors"
)

// Timeout is a Postgres timeout
const Timeout = 5

// Store contains all repositories
type Store struct {
	DB        *pg.DB
	Trademark TrademarkRepository
}

// New creates new store
func New() (*Store, error) {
	pgDB, err := Dial()
	if err != nil {
		return nil, errors.Wrap(err, "store.Dial")
	}

	log.Println("Running PostgreSQL migrations...")
	if err := runPgMigrations(); err != nil {
		return nil, errors.Wrap(err, "store.runPgMigrations")
	}

	store := &Store{pgDB, repositories.NewTrademarkRepository(pgDB)}

	return store, nil
}

// Dial creates new database connection to postgres
func Dial() (*pg.DB, error) {
	cfg := config.Get()
	if cfg.PgURL == "" {
		return nil, nil
	}
	pgOpts, err := pg.ParseURL(cfg.PgURL)
	if err != nil {
		return nil, err
	}

	pgDB := pg.Connect(pgOpts)

	_, err = pgDB.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}

	pgDB.WithTimeout(time.Second * time.Duration(Timeout))

	return pgDB, nil
}
