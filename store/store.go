package store

import (
	"log"
	"time"

	"github.com/denchick/trademarks/config"
	"github.com/denchick/trademarks/store/repositories"
	"github.com/go-pg/pg/v10"
	"github.com/pkg/errors"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
		return nil, errors.Wrap(err, "Postgres dial failed")
	}

	log.Println("Running PostgreSQL migrations...")
	if err := runPgMigrations(); err != nil {
		return nil, errors.Wrap(err, "runPgMigrations failed")
	}

	store := &Store{pgDB, repositories.NewTrademarkRepository(pgDB)}

	return store, nil
}

// KeepAlivePollPeriod is a Pg/MySQL keepalive check time period
const KeepAlivePollPeriod = 3

// KeepAlivePg makes sure PostgreSQL is alive and reconnects if needed
func (store *Store) KeepAlivePg() {
	var err error
	for {
		// Check if PostgreSQL is alive every 3 seconds
		time.Sleep(time.Second * KeepAlivePollPeriod)
		lostConnect := false
		if store.DB == nil {
			lostConnect = true
		} else if _, err = store.DB.Exec("SELECT 1"); err != nil {
			lostConnect = true
		}
		if !lostConnect {
			continue
		}
		log.Println("[store.KeepAlivePg] Lost PostgreSQL connection. Restoring...")
		store.DB, err = Dial()
		if err != nil {
			log.Fatal(err)
			continue
		}
		log.Println("[store.KeepAlivePg] PostgreSQL reconnected")
	}
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
