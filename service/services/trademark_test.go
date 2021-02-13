package services

import (
	"errors"
	"testing"

	"github.com/denchick/trademarks/models"
	"github.com/denchick/trademarks/store"
	"github.com/denchick/trademarks/store/mocks"
	"github.com/stretchr/testify/assert"
)

var trademark = models.DBTrademark{
	ID:                      1,
	ApplicationNumber:       "018004930",
	ApplicationDate:         "2019-01-01",
	RegistrationDate:        "2019-12-04",
	ApplicationLanguageCode: "en",
	SecondLanguageCode:      "it",
	ExpiryDate:              "2029-01-01",
	Name:                    "UNISON",
}

func TestGetTrademarksWhenTrademarkExist(t *testing.T) {
	name := "UNISON"
	repository := &mocks.TrademarkRepository{}
	service := NewTrademarkService(&store.Store{Trademark: repository})
	repository.On("FindByName", name).Return([]*models.DBTrademark{&trademark}, nil)

	expected := trademark.ToTrademark()
	actual, err := service.GetTrademarks(name, false)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(actual))
	assert.ObjectsAreEqual(expected, actual[0])
}

func TestGetTrademarksWhenTrademarkDoesNotExist(t *testing.T) {
	name := "WRONG NAME"
	repository := &mocks.TrademarkRepository{}
	service := NewTrademarkService(&store.Store{Trademark: repository})
	repository.On("FindByName", name).Return(nil, nil)

	actual, err := service.GetTrademarks(name, false)

	assert.NoError(t, err)
	assert.Nil(t, actual)
}

func TestGetTrademarksWithSimilar(t *testing.T) {
	name := "UNISON"
	repository := &mocks.TrademarkRepository{}
	service := NewTrademarkService(&store.Store{Trademark: repository})
	repository.On("FindSimilar", name).Return([]*models.DBTrademark{&trademark}, nil)

	expected := trademark.ToTrademark()
	actual, err := service.GetTrademarks(name, true)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(actual))
	assert.ObjectsAreEqual(expected, actual[0])
}

func TestGetTrademarksWithSimilarWhenError(t *testing.T) {
	name := "UNISON"
	repository := &mocks.TrademarkRepository{}
	service := NewTrademarkService(&store.Store{Trademark: repository})
	repository.On("FindSimilar", name).Return(nil, errors.New("some error"))

	actual, err := service.GetTrademarks(name, true)

	assert.Nil(t, actual)
	assert.Error(t, err, "services.GetTrademarks: some error")
}
