package services

import (
	"strings"
	"github.com/bshome19/Flexera_Code_Test/constants"
	"github.com/bshome19/Flexera_Code_Test/models"
)


func CalculateMinNumberOfCopies(records [][]string) int {
	desktopMap := make(map[string]bool)
	minimumCopies := 0
	duplicateCheck := make(map[string]bool)

	slice1 := []models.ABC{}

	for _, record := range records[1:] {		

		recordsData := models.ABC{

			ComputerID:    record[0],
			UserID:        record[1],
			ApplicationID: record[2],
			ComputerType:  strings.ToUpper(record[3]),
			Comment:       record[4],
		}

		uniqueKey := recordsData.UserID + "_" + recordsData.ComputerID
		if duplicateCheck[uniqueKey] || recordsData.ApplicationID != constants.ApplicationID {
			continue
		} 

		slice1 = append(slice1, recordsData)
		duplicateCheck[uniqueKey] = true

		if recordsData.ComputerType == "DESKTOP" {
			desktopMap[recordsData.UserID] = true
			minimumCopies += 1
		}

	}

	for _, record := range slice1 {
		if record.ComputerType == "LAPTOP" && desktopMap[record.UserID] != true {
				minimumCopies += 1
			}
	}
	return minimumCopies
}