package utils

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCSV(filepath string) [][]string {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV records:", err)
	}
	return records
}
