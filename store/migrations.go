package store

import (
	"github.com/denchick/trademarks/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/pkg/errors"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// runPgMigrations runs Postgres migrations
func runPgMigrations() error {
	cfg := config.Get()
	if cfg.PgMigrationsPath == "" {
		return nil
	}
	if cfg.PgURL == "" {
		return errors.New("No cfg.PgURL provided")
	}
	m, err := migrate.New(
		cfg.PgMigrationsPath,
		cfg.PgURL,
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
