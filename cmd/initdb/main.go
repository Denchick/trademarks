package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/vacuumlabs-interviews/3rd-round-Denis-Volkov/xmlparser"
)

func getXMLsDirectory() string {
	directory := flag.String("directory", "", "Path to directory with XMLs")
	flag.Parse()
	return *directory
}

func main() {
	xmlsDirectory := getXMLsDirectory()
	if len(xmlsDirectory) == 0 {
		fmt.Println("Usage: ./initdb --directory ") // TODO improve usage example
		return
	}
	trademarks := xmlparser.GetTrademarks(xmlsDirectory)
	log.Printf("Parsed %d trademarks successfully\n", len(trademarks))

	csvfile, err := os.Create("output.csv")
	if err != nil {
		log.Fatal(err)
		return
	}

	defer csvfile.Close()
	writer := csv.NewWriter(csvfile)
	writer.Write([]string{"ApplicationNumber", "ApplicationDate", "RegistrationDate", "ApplicationLanguageCode", "SecondLanguageCode", "ExpiryDate", "Name"})
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
