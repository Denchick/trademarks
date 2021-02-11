package models

// Trademark is trademark model
type Trademark struct {
	ApplicationNumber       string `json:"application_number"`
	ApplicationDate         string `json:"application_date"`
	RegistrationDate        string `json:"registration_date"`
	ApplicationLanguageCode string `json:"application_language_code"`
	SecondLanguageCode      string `json:"second_language_code"`
	ExpiryDate              string `json:"expiry_date"`
	Name                    string `json:"name"`
}

// DBTrademark is database representation of trademark model
type DBTrademark struct {
	ID                      uint
	ApplicationNumber       string `gorm:"unique"`
	ApplicationDate         string
	RegistrationDate        string
	ApplicationLanguageCode string
	SecondLanguageCode      string
	ExpiryDate              string
	Name                    string `gorm:"index"`
}

// TableName overrides default table name for gorm
func (DBTrademark) TableName() string {
	return "trademarks"
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
