package csv

import (
	"os"
	"GoStori/models"
	"github.com/gocarina/gocsv"
)

func ProcessCSVFile(filePath string) ([]models.Transaction, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var transactions []models.Transaction
	if err := gocsv.UnmarshalFile(file, &transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}