package main

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/vacuumlabs-interviews/3rd-round-Denis-Volkov/logger"
	"github.com/vacuumlabs-interviews/3rd-round-Denis-Volkov/models"
)

type xmlRoot struct {
	Transaction  xml.Name     `xml:"Transaction"`
	XMLTrademark xmlTrademark `xml:"TradeMarkTransactionBody>TransactionContentDetails>TransactionData>TradeMarkDetails>TradeMark"`
}

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

// XMLParser ...
type XMLParser struct {
	logger *logger.Logger
}

// NewXMLParser ...
func NewXMLParser(logger *logger.Logger) *XMLParser {
	return &XMLParser{logger}
}

func (parser *XMLParser) parseXML(filepath string) (*xmlTrademark, error) {
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
	return &root.XMLTrademark, nil
}

func (parser *XMLParser) getXMLPaths(rootpath string) ([]string, error) {
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
	return xmlPaths, err
}

// GetWordTrademarks ...
func (parser *XMLParser) GetWordTrademarks(directory string) []*models.Trademark {
	xmlPaths, err := parser.getXMLPaths(directory)
	if err != nil {
		parser.logger.Err(err)
		return make([]*models.Trademark, 0)
	}

	parser.logger.Debug().Msgf("Found %d XMLs", len(xmlPaths))
	trademarks := make([]*models.Trademark, 0, len(xmlPaths))
	for _, pathToXML := range xmlPaths {
		xmlTrademark, err := parser.parseXML(pathToXML)
		if err != nil {
			parser.logger.Err(err)
			continue
		}

		if xmlTrademark.OperationCode == "Insert" && xmlTrademark.MarkFeature == "Word" && xmlTrademark.MarkCurrentStatusCode == "Registered" {
			trademarks = append(trademarks, xmlTrademark.toTrademark())
			parser.logger.Debug().Msgf("%v", xmlTrademark)
		}
	}
	return trademarks
}
