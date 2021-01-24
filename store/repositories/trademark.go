package repositories

import (
	"context"

	"github.com/vacuumlabs-interviews/3rd-round-Denis-Volkov/models"
	"gorm.io/gorm"
)

// TrademarkRepository ...
type TrademarkRepository struct {
	db  *gorm.DB
}

// NewTrademarkRepository ...
func NewTrademarkRepository(db *gorm.DB) *TrademarkRepository {
	return &TrademarkRepository{db: db}
}

// FindTrademarkByName retrieves trademark from DB
func (repository *TrademarkRepository) FindTrademarkByName(ctx context.Context, name string) (*models.DBTrademark, error) {
	var trademark models.DBTrademark
	result := repository.db.Where("name = ?", name).First(&trademark)
	if result.Error != nil {
		return nil, result.Error
	}
	return &trademark, nil
}

// FuzzyFindTrademarkByName retrieves  similar trademark from DB
func (repository *TrademarkRepository) FuzzyFindTrademarkByName(ctx context.Context, name string) (*models.DBTrademark, error) {
	return repository.FindTrademarkByName(ctx, name) // TODO implement
}