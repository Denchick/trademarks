package models

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
