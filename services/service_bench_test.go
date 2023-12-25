package services

import (
	"testing"

)

func BenchmarkCalculateMinNumberOfCopies(b *testing.B) {
	// Test case 1: One copy is required
	input1 := [][]string{
		{"1", "1", "374", "LAPTOP", "Exported from System A"},
		{"2", "1", "374", "DESKTOP", "Exported from System A"},
	}
	for i := 0; i < b.N; i++ {
		_ = CalculateMinNumberOfCopies(input1)
	}	

	// Test case 2: Three copies are required
	input2 := [][]string{
		{"1", "1", "374", "LAPTOP", "Exported from System A"},
		{"2", "1", "374", "DESKTOP", "Exported from System A"},
		{"3", "2", "374", "DESKTOP", "Exported from System A"},
		{"4", "2", "374", "DESKTOP", "Exported from System A"},
	}
	for i := 0; i < b.N; i++ {
		_ = CalculateMinNumberOfCopies(input2)
	}	

	// Test case 3: Two copies are required (handling duplicates)
	input3 := [][]string{
		{"1", "1", "374", "LAPTOP", "Exported from System A"},
		{"2", "2", "374", "DESKTOP", "Exported from System A"},
		{"3", "2", "374", "desktop", "Exported from System B"},
	}
	for i := 0; i < b.N; i++ {
		_ = CalculateMinNumberOfCopies(input3)
	}	
}
