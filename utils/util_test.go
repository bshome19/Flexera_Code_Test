package utils

import (
	"os"
	"path/filepath"
	"testing"
	"github.com/bshome19/Flexera_Code_Test/constants"
)

func TestReadCSV(t *testing.T) {
	currentDir, _ := os.Getwd()
	// Test case 1: 
	smallFilePath := filepath.Join(currentDir, "/../", constants.SmallFilePath)
	result1 := ReadCSV(smallFilePath)
	if len(result1) != 220131 {
		t.Errorf("Test case 1 failed: Expected 220131, got %d", len(result1))
	}

	// Test case 2:
	largeFilePath := filepath.Join(currentDir, "/../", constants.LargeFilePath)
	result2 := ReadCSV(largeFilePath)
	if len(result2) != 21998526 {
		t.Errorf("Test case 2 failed: Expected 21998526, got %d", len(result2))
	}
}