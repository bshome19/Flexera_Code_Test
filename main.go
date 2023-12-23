package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"github.com/bshome19/Flexera_Code_Test/utils"
)
 
func main() {
	file, err := os.Open("sample-small.csv")
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV records:", err)
		return
	}
	minimumCopiesRequired := utils.CalculateMinNumberOfCopies(records)
	
	fmt.Println(minimumCopiesRequired)
}
