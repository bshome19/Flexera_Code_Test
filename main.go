package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bshome19/Flexera_Code_Test/constants"
	"github.com/bshome19/Flexera_Code_Test/services"
	"github.com/bshome19/Flexera_Code_Test/utils"
)
 
func main() {

	currentDir, _ := os.Getwd()
	filePath := filepath.Join(currentDir, constants.SmallFilePath)


	records := utils.ReadCSV(filePath)
	minimumCopiesRequired := services.CalculateMinNumberOfCopies(records)

	fmt.Println(minimumCopiesRequired)
}
