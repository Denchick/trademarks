package services

import (
	"github.com/denchick/trademarks/models"
	"github.com/denchick/trademarks/store"
	"github.com/pkg/errors"
)

// TrademarkService ...
type TrademarkService struct {
	store *store.Store
}

// NewTrademarkService ...
func NewTrademarkService(store *store.Store) *TrademarkService {
	return &TrademarkService{store}
}

// GetTrademarks ...
func (service *TrademarkService) GetTrademarks(name string, similar bool) ([]*models.Trademark, error) {
	trademarks, err := service.getTrademarks(name, similar)
	if err != nil {
		return nil, errors.Wrap(err, "services.GetTrademarks")
	}
	return service.toWeb(trademarks), nil
}


func (service *TrademarkService) getTrademarks(name string, similar bool) ([]*models.DBTrademark, error) {
	if similar {
		return service.store.Trademark.FindSimilar(name)
	}
	return service.store.Trademark.FindByName(name)
}

func (service *TrademarkService) toWeb(oldTrademarks []*models.DBTrademark) []*models.Trademark {
	var newTrademarks []*models.Trademark
	for _, trademark := range oldTrademarks {
		newTrademarks = append(newTrademarks, trademark.ToTrademark())
	}
	return newTrademarks
}
