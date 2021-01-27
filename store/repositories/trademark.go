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
	// https://www.postgresql.org/docs/13/pgtrgm.html#id-1.11.7.40.8
	db.Exec("CREATE EXTENSION pg_trgm;")
	return &TrademarkRepository{db: db}
}

// FindTrademarkByName retrieves trademark from DB
func (repository *TrademarkRepository) FindTrademarkByName(ctx context.Context, name string) (*models.DBTrademark, error) {
	var trademark models.DBTrademark
	result := repository.db.Where("name = ?", name).First(&trademark)

	if result.Error == nil {
		return &trademark, nil
	}
	return nil, result.Error
}

// FindSimilarTrademarks retrieves similar trademarks from DB
func (repository *TrademarkRepository) FindSimilarTrademarks(ctx context.Context, name string) ([]*models.DBTrademark, error) {
	var trademarks []*models.DBTrademark
	// I can't build a query using gorm.Clause https://gorm.io/docs/query.html#Order
	result := repository.db.Raw("SELECT * FROM db_trademarks ORDER BY @trademark <-> name LIMIT 3;", sql.Named("trademark", name)).Scan(&trademarks)
	if result.Error == nil {
		return trademarks, nil
	}
	return nil, result.Error
}
