package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

// rename to xml2sql
func getXMLsDirectory() string {
	directory := flag.String("directory", "", "Path to directory with XMLs")
	flag.Parse()
	return *directory
}

func main() {
	xmlsDirectory := getXMLsDirectory()
	if len(xmlsDirectory) == 0 {
		fmt.Println("Generates csv file with information about trademarks for importing into database.")
		fmt.Println("Usage: ./xml2csv --directory /path/to/xmls")
		return
	}
	trademarks := getTrademarks(xmlsDirectory)
	log.Printf("Parsed %d trademarks successfully\n", len(trademarks))

	csvfile, err := os.Create("output.csv")
	if err != nil {
		log.Fatal(err)
		return
	}

	defer csvfile.Close()
	writer := csv.NewWriter(csvfile)
	writer.Write([]string{"application_number", "application_date", "registration_date", "application_language_code", "second_language_code", "expiry_date", "name"})
	for _, trademark := range trademarks {
		record := []string{trademark.ApplicationNumber, trademark.ApplicationDate, trademark.RegistrationDate, trademark.ApplicationLanguageCode, trademark.SecondLanguageCode, trademark.ExpiryDate, trademark.Name}
		err := writer.Write(record)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
	writer.Flush()
}
