package main

import (
	"encoding/csv"
	"os"

	"github.com/denchick/trademarks/logger"
	"github.com/denchick/trademarks/models"
)

// XML2CSVWriter ...
type XML2CSVWriter struct {
	logger *logger.Logger
}

// NewXML2CSVWriter creates new XML2CSVWriter
func NewXML2CSVWriter(logger *logger.Logger) *XML2CSVWriter {
	return &XML2CSVWriter{logger}
}

// WriteToFile writes trademarks to file
func (xml2csvWriter *XML2CSVWriter) WriteToFile(filename string, trademarks []*models.Trademark) {
	csvfile, err := os.Create(filename)
	if err != nil {
		xml2csvWriter.logger.Err(err)
		return
	}

	defer csvfile.Close()
	writer := csv.NewWriter(csvfile)
	writer.Comma = ';'
	if err != nil {
		xml2csvWriter.logger.Err(err)
		return
	}

	writer.Write([]string{"application_number", "application_date", "registration_date", "application_language_code", "second_language_code", "expiry_date", "name"})
	for _, trademark := range trademarks {
		record := []string{trademark.ApplicationNumber, trademark.ApplicationDate, trademark.RegistrationDate, trademark.ApplicationLanguageCode, trademark.SecondLanguageCode, trademark.ExpiryDate, trademark.Name}
		err := writer.Write(record)
		if err != nil {
			xml2csvWriter.logger.Err(err)
			return
		}
	}
	writer.Flush()
}
