package xmlparser

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// TradeMark represents trade mark
type TradeMark struct {
	TradeMark                     xml.Name `xml:"Transaction>TradeMarkTransactionBody>TransactionContentDetails>TransactionData>TradeMarkDetails>TradeMark"`
	OperationCode                 string   `xml:"operationCode,attr"`
	RegistrationOfficeCode        string   `xml:"RegistrationOfficeCode"`
	ApplicationNumber             string   `xml:"ApplicationNumber"`
	ApplicationDate               string   `xml:"ApplicationDate"`
	RegistrationDate              string   `xml:"RegistrationDate"`
	ApplicationLanguageCode       string   `xml:"ApplicationLanguageCode"`
	SecondLanguageCode            string   `xml:"SecondLanguageCode"`
	ExpiryDate                    string   `xml:"ExpiryDate"`
	KindMark                      string   `xml:"KindMark"`
	MarkFeature                   string   `xml:"MarkFeature"`
	TradeDistinctivenessIndicator string   `xml:"TradeDistinctivenessIndicator"`
}

func parseXML(filepath string) (*TradeMark, error) {
	xmlFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err) // TODO use logger
		return nil, err
	}
	fmt.Println("Successfully Opened " + filepath)
	defer xmlFile.Close()
	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println(err) // TODO use logger
		return nil, err
	}

	var tradeMark TradeMark
	xml.Unmarshal(byteValue, &tradeMark)
	return &tradeMark, nil
}

func getXMLPaths(rootpath string) []string { // TODO add error handling
	xmlPaths := make([]string, 0, 10) // TODO why 10?
	err := filepath.Walk(rootpath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".xml" {
			xmlPaths = append(xmlPaths, path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("walk error [%v]\n", err)
	}
	return xmlPaths
}

// ParseXmls parsing trademarks from XMLs in directory
func ParseXmls(directory string) []TradeMark {
	tradeMarks := make([]TradeMark, 0, 0)
	for _, pathToXML := range getXMLPaths(directory) {
		tradeMark, err := parseXML(pathToXML)
		if err != nil {
			fmt.Println(err)
			continue
		}
		tradeMarks = append(tradeMarks, *tradeMark)
	}
	return tradeMarks
}
