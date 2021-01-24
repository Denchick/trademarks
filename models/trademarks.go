package models

import (
	"encoding/xml"

	"gorm.io/gorm"
)

// Trademark is trademark model
type Trademark struct {
	ApplicationNumber       string
	ApplicationDate         string
	RegistrationDate        string
	ApplicationLanguageCode string
	SecondLanguageCode      string
	ExpiryDate              string
	Name                    string
}

// XMLRoot helps to parse XML tree
type XMLRoot struct {
	Transaction  xml.Name     `xml:"Transaction"`
	XMLTrademark XMLTrademark `xml:"TradeMarkTransactionBody>TransactionContentDetails>TransactionData>TradeMarkDetails>TradeMark"`
}

// XMLTrademark is XML representation of trademark model
type XMLTrademark struct {
	Trademark               xml.Name `xml:"TradeMark"`
	OperationCode           string   `xml:"operationCode,attr"`
	RegistrationOfficeCode  string   `xml:"RegistrationOfficeCode"`
	ApplicationNumber       string   `xml:"ApplicationNumber"`
	ApplicationDate         string   `xml:"ApplicationDate"`
	RegistrationDate        string   `xml:"RegistrationDate"`
	ApplicationLanguageCode string   `xml:"ApplicationLanguageCode"`
	MarkCurrentStatusCode   string   `xml:"MarkCurrentStatusCode"`
	SecondLanguageCode      string   `xml:"SecondLanguageCode"`
	ExpiryDate              string   `xml:"ExpiryDate"`
	MarkFeature             string   `xml:"MarkFeature"`
	Name                    string   `xml:"WordMarkSpecification>MarkVerbalElementText"`
}

// DBTrademark is database representation of trademark model
type DBTrademark struct {
	gorm.Model
	ApplicationNumber       string `gorm:"unique"`
	ApplicationDate         string
	RegistrationDate        string
	ApplicationLanguageCode string
	SecondLanguageCode      string
	ExpiryDate              string
	Name                    string `gorm:"index"`
}

// ToTrademark converts XMLTrademark to Trademark
func (trademark *XMLTrademark) ToTrademark() *Trademark {
	return &Trademark{
		ApplicationNumber:       trademark.ApplicationNumber,
		ApplicationDate:         trademark.ApplicationDate,
		RegistrationDate:        trademark.RegistrationDate,
		ApplicationLanguageCode: trademark.ApplicationLanguageCode,
		SecondLanguageCode:      trademark.SecondLanguageCode,
		ExpiryDate:              trademark.ExpiryDate,
		Name:                    trademark.Name,
	}
}

// ToTrademark converts DBTrademark to Trademark
func (trademark *DBTrademark) ToTrademark() *Trademark {
	return &Trademark{
		ApplicationNumber:       trademark.ApplicationNumber,
		ApplicationDate:         trademark.ApplicationDate,
		RegistrationDate:        trademark.RegistrationDate,
		ApplicationLanguageCode: trademark.ApplicationLanguageCode,
		SecondLanguageCode:      trademark.SecondLanguageCode,
		ExpiryDate:              trademark.ExpiryDate,
		Name:                    trademark.Name,
	}
}

// ToDB converts Trademark to DBTrademark
func (trademark *Trademark) ToDB() *DBTrademark {
	return &DBTrademark{
		ApplicationNumber:       trademark.ApplicationNumber,
		ApplicationDate:         trademark.ApplicationDate,
		RegistrationDate:        trademark.RegistrationDate,
		ApplicationLanguageCode: trademark.ApplicationLanguageCode,
		SecondLanguageCode:      trademark.SecondLanguageCode,
		ExpiryDate:              trademark.ExpiryDate,
		Name:                    trademark.Name,
	}
}
