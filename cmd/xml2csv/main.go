package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"

	"github.com/vacuumlabs-interviews/3rd-round-Denis-Volkov/logger"
)

func getParams() (directory string, logLevel string) {
	directoryFlag := flag.String("directory", "", "Path to directory with XMLs")
	verbose := flag.Bool("verbose", false, "Turn on verbose output")
	flag.Parse()
	directory = *directoryFlag
	logLevel = "error"
	if *verbose {
		logLevel = "debug"
	}
	return
}

func createCsvWriter() (*csv.Writer, error) {
	csvfile, err := os.Create("trademarks.csv")
	if err != nil {
		return nil, err
	}

	defer csvfile.Close() // TODO dangerous place
	writer := csv.NewWriter(csvfile)
	writer.Comma = ';'
	return writer, nil
}

func main() {
	pathToXmls, logLevel := getParams()
	logger := logger.Get(logLevel)
	parser := NewXMLParser(logger)
	if len(pathToXmls) == 0 {
		fmt.Println("Generates csv file with information about trademarks for importing into database.")
		fmt.Println("Usage: ./xml2csv --directory /path/to/xmls --verbose")
		return
	}
	trademarks := parser.GetWordTrademarks(pathToXmls)
	parser.logger.Debug().Msgf("Parsed %d trademarks successfully", len(trademarks))
	writer, err := createCsvWriter()
	if err != nil {
		logger.Err(err)
		return
	}

	writer.Write([]string{"application_number", "application_date", "registration_date", "application_language_code", "second_language_code", "expiry_date", "name"})
	for i, trademark := range trademarks {
		record := []string{fmt.Sprint(i), trademark.ApplicationNumber, trademark.ApplicationDate, trademark.RegistrationDate, trademark.ApplicationLanguageCode, trademark.SecondLanguageCode, trademark.ExpiryDate, trademark.Name}
		err := writer.Write(record)
		if err != nil {
			logger.Err(err)
			return
		}
	}
	writer.Flush()
}
