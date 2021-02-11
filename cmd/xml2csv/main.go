package main

import (
	"flag"
	"fmt"

	"github.com/denchick/trademarks/logger"
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

func main() {
	pathToXmls, logLevel := getParams()
	logger := logger.Get(logLevel)
	xmlParser := NewXMLParser(logger)
	csvWriter := NewXML2CSVWriter(logger)
	if len(pathToXmls) == 0 {
		fmt.Println("Generates csv file with information about trademarks for importing into database.")
		fmt.Println("Usage: ./xml2csv --directory /path/to/xmls --verbose")
		return
	}

	trademarks := xmlParser.GetWordTrademarks(pathToXmls)
	csvWriter.WriteToFile("trademarks.csv", trademarks)
}
