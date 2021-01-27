package store

import (
	"github.com/pkg/errors"
	"github.com/vacuumlabs-interviews/3rd-round-Denis-Volkov/config"
	"github.com/vacuumlabs-interviews/3rd-round-Denis-Volkov/models"
	"github.com/vacuumlabs-interviews/3rd-round-Denis-Volkov/store/repositories"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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

	return &store, nil
}
