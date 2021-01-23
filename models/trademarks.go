package models

import "encoding/xml"

// TradeMark is trademark model
type TradeMark struct {
	RegistrationOfficeCode  string
	ApplicationNumber       string
	ApplicationDate         string
	RegistrationDate        string
	ApplicationLanguageCode string
	SecondLanguageCode      string
	ExpiryDate              string
}

// XMLTradeMark is XML representation of trademark model
type XMLTradeMark struct {
	TradeMark               xml.Name `xml:"Transaction>TradeMarkTransactionBody>TransactionContentDetails>TransactionData>TradeMarkDetails>TradeMark"`
	OperationCode           string   `xml:"operationCode,attr"`
	RegistrationOfficeCode  string   `xml:"RegistrationOfficeCode"`
	ApplicationNumber       string   `xml:"ApplicationNumber"`
	ApplicationDate         string   `xml:"ApplicationDate"`
	RegistrationDate        string   `xml:"RegistrationDate"`
	ApplicationLanguageCode string   `xml:"ApplicationLanguageCode"`
	SecondLanguageCode      string   `xml:"SecondLanguageCode"`
	ExpiryDate              string   `xml:"ExpiryDate"`
	MarkFeature             string   `xml:"MarkFeature"`
}

// ToTradeMark converts XMLTradeMark to TradeMark
func (trademark *XMLTradeMark) ToTradeMark() *TradeMark {
	return &TradeMark{
		RegistrationOfficeCode: trademark.RegistrationOfficeCode,
		ApplicationNumber: trademark.ApplicationNumber,
		ApplicationDate: trademark.ApplicationDate,
		RegistrationDate: trademark.RegistrationDate,
		ApplicationLanguageCode: trademark.ApplicationLanguageCode,
		SecondLanguageCode: trademark.SecondLanguageCode,
		ExpiryDate: trademark.ExpiryDate,
	}
}
