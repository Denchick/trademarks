package repositories

import (
	"github.com/denchick/trademarks/models"
	"github.com/go-pg/pg/v10"
	"github.com/pkg/errors"
)

// TrademarkRepository ...
type TrademarkRepository struct {
	db *pg.DB
}

// NewTrademarkRepository ...
func NewTrademarkRepository(db *pg.DB) *TrademarkRepository {
	return &TrademarkRepository{db}
}

// FindByName retrieves trademark from DB
func (repository *TrademarkRepository) FindByName(name string) ([]*models.DBTrademark, error) {
	var trademarks []*models.DBTrademark
	err := repository.db.Model(&trademarks).
		Where("name = ?", name).
		Select()
	if err != nil && err != pg.ErrNoRows {
		return nil, errors.Wrap(err, "store.repositories.FindByName")
	}
	return trademarks, nil
}

// FindSimilar retrieves similar trademarks from DB
func (repository *TrademarkRepository) FindSimilar(name string) ([]*models.DBTrademark, error) {
	var trademarks []*models.DBTrademark
	err := repository.db.Model(&trademarks).
		OrderExpr("? <-> name", name).
		Limit(3).
		Select()
	if err != nil && err != pg.ErrNoRows {
		return nil, errors.Wrap(err, "store.repositories.FindSimilar")
	}
	return trademarks, nil
}
