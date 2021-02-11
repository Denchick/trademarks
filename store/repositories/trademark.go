package repositories

import (
	"github.com/denchick/trademarks/models"
	"github.com/go-pg/pg"
)

// TrademarkRepository ...
type TrademarkRepository struct {
	db *pg.DB
}

// NewTrademarkRepository ...
func NewTrademarkRepository(db *pg.DB) *TrademarkRepository {
	// https://www.postgresql.org/docs/13/pgtrgm.html#id-1.11.7.40.8
	db.Exec("CREATE EXTENSION pg_trgm;")
	return &TrademarkRepository{db}
}

// FindTrademarkByName retrieves trademark from DB
func (repository *TrademarkRepository) FindTrademarkByName(name string) (*models.DBTrademark, error) {
	trademark := &models.DBTrademark{}
	err := repository.db.Model(trademark).
		Where("name = ?", name).
		First()
	return trademark, err
}

// FindSimilarTrademarks retrieves similar trademarks from DB
func (repository *TrademarkRepository) FindSimilarTrademarks(name string) ([]*models.DBTrademark, error) {
	var trademarks []*models.DBTrademark
	err := repository.db.Model(trademarks).
		OrderExpr("? <-> name", name).
		Limit(3).
		Select()
	return trademarks, err
}
