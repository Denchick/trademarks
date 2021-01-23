package xmlparser

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/vacuumlabs-interviews/3rd-round-Denis-Volkov/models"
)

type xmlRoot struct {
	Transaction  xml.Name            `xml:"Transaction"`
	XMLTradeMark models.XMLTradeMark `xml:"TradeMarkTransactionBody>TransactionContentDetails>TransactionData>TradeMarkDetails>TradeMark"`
}

func parseXML(filepath string) (*models.XMLTradeMark, error) {
	xmlFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err) // TODO use logger
		return nil, err
	}
	defer xmlFile.Close()
	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println(err) // TODO use logger
		return nil, err
	}

	var xmlRoot xmlRoot
	xml.Unmarshal(byteValue, &xmlRoot)
	fmt.Println(xmlRoot)
	return &xmlRoot.XMLTradeMark, nil
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

// GetTradeMarks parsing trademarks from XMLs in directory
func GetTradeMarks(directory string) []models.TradeMark {
	tradeMarks := make([]models.TradeMark, 0, 0)
	for _, pathToXML := range getXMLPaths(directory) {
		xmlTradeMark, err := parseXML(pathToXML)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if xmlTradeMark.OperationCode == "Insert" && xmlTradeMark.MarkFeature == "Word" {
			tradeMarks = append(tradeMarks, *xmlTradeMark.ToTradeMark())
		}
	}
	return tradeMarks
}
