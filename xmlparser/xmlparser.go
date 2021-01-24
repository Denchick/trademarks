package xmlparser

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/vacuumlabs-interviews/3rd-round-Denis-Volkov/models"
)

// TODO move to initdb
func parseXML(filepath string) (*models.XMLTrademark, error) {
	xmlFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()
	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return nil, err
	}

	var xmlRoot models.XMLRoot
	xml.Unmarshal(byteValue, &xmlRoot)
	fmt.Println(xmlRoot)
	return &xmlRoot.XMLTrademark, nil
}

func getXMLPaths(rootpath string) ([]string, error) {
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
		return nil, err
	}
	return xmlPaths, nil
}

// GetTrademarks parsing trademarks from XMLs in directory
func GetTrademarks(directory string) []models.Trademark {
	xmlPaths, err := getXMLPaths(directory)
	if err != nil {
		log.Fatal(err)
		return make([]models.Trademark, 0)
	}

	trademarks := make([]models.Trademark, 0, len(xmlPaths))
	for _, pathToXML := range xmlPaths {
		xmlTrademark, err := parseXML(pathToXML)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// TODO extract business logic from here
		if xmlTrademark.OperationCode == "Insert" && xmlTrademark.MarkFeature == "Word" && xmlTrademark.MarkCurrentStatusCode == "Registered" {
			trademarks = append(trademarks, *xmlTrademark.ToTrademark())
		}
	}
	return trademarks
}
