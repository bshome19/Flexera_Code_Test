package utils

import (
	"os"
	"path/filepath"
	"testing"
	"github.com/bshome19/Flexera_Code_Test/constants"
)

func BenchmarkReadCSV(b *testing.B) {
	currentDir, _ := os.Getwd()
	// Test case 1: 
	smallFilePath := filepath.Join(currentDir, "/../", constants.SmallFilePath)
	for i := 0; i < b.N; i++ {
		_ = ReadCSV(smallFilePath)
	}

	// Test case 2:
	largeFilePath := filepath.Join(currentDir, "/../", constants.LargeFilePath)
	for i := 0; i < b.N; i++ {
		_ = ReadCSV(largeFilePath)
	}	
}