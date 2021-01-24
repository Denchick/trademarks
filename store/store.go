package store

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vacuumlabs-interviews/3rd-round-Denis-Volkov/models"
	"github.com/vacuumlabs-interviews/3rd-round-Denis-Volkov/store/repositories"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TrademarkRepository is a store for trademarks
type TrademarkRepository interface {
	FindTrademarkByName(context.Context, string) (*models.DBTrademark, error)
	FuzzyFindTrademarkByName(context.Context, string) (*models.DBTrademark, error)
}

// Store contains all repositories
type Store struct {
	DB        *gorm.DB
	Trademark TrademarkRepository
}

// New creates new store
func New() (*Store, error) {
	dsn := "host=localhost port=7777 user=volkov dbname=trademarksdb password=trademarks"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, errors.Wrap(err, "database connection failed")
	}
	var store Store
	if db != nil {
		store.DB = db
		store.Trademark = repositories.NewTrademarkRepository(db)
	}

	return &store, nil
}
