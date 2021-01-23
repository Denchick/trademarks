package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"

	"github.com/vacuumlabs-interviews/3rd-round-Denis-Volkov/xmlparser"
)

func main() {
	directory := flag.String("directory", "", "Path to directory with XMLs")
	flag.Parse()
	if len(*directory) == 0 {
		fmt.Println("Usage: ./initdb --directory ") // TODO improve usage example
		return
	}

	tradeMarks := xmlparser.GetTradeMarks(*directory)
	fmt.Printf("Parsed %d trademarks successfully\n", len(tradeMarks))

	csvfile, err := os.Create("output.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer csvfile.Close()

	writer := csv.NewWriter(csvfile)
	writer.Write([]string{"RegistrationOfficeCode", "ApplicationNumber", "ApplicationDate", "RegistrationDate", "ApplicationLanguageCode", "SecondLanguageCode", "ExpiryDate"})
	for _, tradeMark := range tradeMarks {
		record := []string{tradeMark.RegistrationOfficeCode, tradeMark.ApplicationNumber, tradeMark.ApplicationDate, tradeMark.RegistrationDate, tradeMark.ApplicationLanguageCode, tradeMark.SecondLanguageCode, tradeMark.ExpiryDate}
		err := writer.Write(record)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
	writer.Flush()
}
