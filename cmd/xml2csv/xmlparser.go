package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/vacuumlabs-interviews/3rd-round-Denis-Volkov/models"
)

// xmlRoot helps to parse XML tree
type xmlRoot struct {
	Transaction  xml.Name     `xml:"Transaction"`
	XMLTrademark xmlTrademark `xml:"TradeMarkTransactionBody>TransactionContentDetails>TransactionData>TradeMarkDetails>TradeMark"`
}

// xmlTrademark is XML representation of trademark model
type xmlTrademark struct {
	Trademark               xml.Name `xml:"TradeMark"`
	OperationCode           string   `xml:"operationCode,attr"`
	RegistrationOfficeCode  string   `xml:"RegistrationOfficeCode"`
	ApplicationNumber       string   `xml:"ApplicationNumber"`
	ApplicationDate         string   `xml:"ApplicationDate"`
	RegistrationDate        string   `xml:"RegistrationDate"`
	ApplicationLanguageCode string   `xml:"ApplicationLanguageCode"`
	MarkCurrentStatusCode   string   `xml:"MarkCurrentStatusCode"`
	SecondLanguageCode      string   `xml:"SecondLanguageCode"`
	ExpiryDate              string   `xml:"ExpiryDate"`
	MarkFeature             string   `xml:"MarkFeature"`
	Name                    string   `xml:"WordMarkSpecification>MarkVerbalElementText"`
}

// toTrademark converts XMLTrademark to Trademark
func (trademark *xmlTrademark) toTrademark() *models.Trademark {
	return &models.Trademark{
		ApplicationNumber:       trademark.ApplicationNumber,
		ApplicationDate:         trademark.ApplicationDate,
		RegistrationDate:        trademark.RegistrationDate,
		ApplicationLanguageCode: trademark.ApplicationLanguageCode,
		SecondLanguageCode:      trademark.SecondLanguageCode,
		ExpiryDate:              trademark.ExpiryDate,
		Name:                    trademark.Name,
	}
}

func parseXML(filepath string) (*xmlTrademark, error) {
	xmlFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()
	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return nil, err
	}

	var root xmlRoot
	xml.Unmarshal(byteValue, &root)
	fmt.Println(root)
	return &root.XMLTrademark, nil
}

func getXMLPaths(rootpath string) ([]string, error) {
	var xmlPaths []string
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
		return nil, err
	}
	return xmlPaths, nil
}

func getTrademarks(directory string) []*models.Trademark {
	xmlPaths, err := getXMLPaths(directory)
	if err != nil {
		log.Fatal(err)
		return make([]*models.Trademark, 0)
	}

	trademarks := make([]*models.Trademark, 0, len(xmlPaths))
	for _, pathToXML := range xmlPaths {
		xmlTrademark, err := parseXML(pathToXML)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// TODO extract business logic from here
		if xmlTrademark.OperationCode == "Insert" && xmlTrademark.MarkFeature == "Word" && xmlTrademark.MarkCurrentStatusCode == "Registered" {
			trademarks = append(trademarks, xmlTrademark.toTrademark())
		}
	}
	return trademarks
}
