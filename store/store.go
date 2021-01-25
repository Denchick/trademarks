package store

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vacuumlabs-interviews/3rd-round-Denis-Volkov/config"
	"github.com/vacuumlabs-interviews/3rd-round-Denis-Volkov/models"
	"github.com/vacuumlabs-interviews/3rd-round-Denis-Volkov/store/repositories"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TrademarkRepository is a store for trademarks
type TrademarkRepository interface {
	FindTrademarkByName(context.Context, string, bool) (*models.DBTrademark, error)
}

// Store contains all repositories
type Store struct {
	DB        *gorm.DB
	Trademark TrademarkRepository
}

// New creates new store
func New() (*Store, error) {
	cfg := config.Get()
	db, err := gorm.Open(postgres.Open(cfg.PgURL), &gorm.Config{})

	if err != nil {
		return nil, errors.Wrap(err, "database connection failed")
	}
	var store Store = Store{
		DB:        db,
		Trademark: repositories.NewTrademarkRepository(db),
	}
	db.AutoMigrate(&models.DBTrademark{})
	db.Raw("CREATE EXTENSION pg_trgm;")

	return &store, nil
}
