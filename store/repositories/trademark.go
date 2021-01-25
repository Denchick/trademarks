package repositories

import (
	"context"
	"database/sql"

	"github.com/vacuumlabs-interviews/3rd-round-Denis-Volkov/models"
	"gorm.io/gorm"
)

// TrademarkRepository ...
type TrademarkRepository struct {
	db *gorm.DB
}

// NewTrademarkRepository ...
func NewTrademarkRepository(db *gorm.DB) *TrademarkRepository {
	return &TrademarkRepository{db: db}
}

// FindTrademarkByName retrieves trademark from DB
func (repository *TrademarkRepository) FindTrademarkByName(ctx context.Context, name string) (*models.DBTrademark, error) {
	var trademark models.DBTrademark
	result := repository.db.Where("name = ?", name).First(&trademark)
	if result.Error != nil { // TODO use https://gorm.io/docs/error_handling.html#ErrRecordNotFound
		return nil, result.Error
	}
	return &trademark, nil
}

// FindSimilarTrademarks retrieves similar trademarks from DB
func (repository *TrademarkRepository) FindSimilarTrademarks(ctx context.Context, name string) ([]*models.DBTrademark, error) {
	var trademarks []*models.DBTrademark
	result := repository.db.Raw("SELECT * FROM db_trademarks AS a ORDER BY @trademark <-> a.name limit 3;", sql.Named("trademark", name)).Scan(&trademarks)
	
	if result.Error != nil { // TODO use https://gorm.io/docs/error_handling.html#ErrRecordNotFound
		return nil, result.Error
	}
	return trademarks, nil
}
