package services

import (
	"testing"

)

func TestCalculateMinNumberOfCopies(t *testing.T) {
	// Test case 1: One copy is required
	input1 := [][]string{
		{"1", "1", "374", "LAPTOP", "Exported from System A"},
		{"2", "1", "374", "DESKTOP", "Exported from System A"},
	}
	result1 := CalculateMinNumberOfCopies(input1)
	if result1 != 1 {
		t.Errorf("Test case 1 failed: Expected 1 copy, got %d", result1)
	}

	// Test case 2: Three copies are required
	input2 := [][]string{
		{"1", "1", "374", "LAPTOP", "Exported from System A"},
		{"2", "1", "374", "DESKTOP", "Exported from System A"},
		{"3", "2", "374", "DESKTOP", "Exported from System A"},
		{"4", "2", "374", "DESKTOP", "Exported from System A"},
	}
	result2 := CalculateMinNumberOfCopies(input2)
	if result2 != 3 {
		t.Errorf("Test case 2 failed: Expected 3 copies, got %d", result2)
	}

	// Test case 3: Two copies are required (handling duplicates)
	input3 := [][]string{
		{"1", "1", "374", "LAPTOP", "Exported from System A"},
		{"2", "2", "374", "DESKTOP", "Exported from System A"},
		{"3", "2", "374", "desktop", "Exported from System B"},
	}
	result3 := CalculateMinNumberOfCopies(input3)
	if result3 != 2 {
		t.Errorf("Test case 3 failed: Expected 2 copies, got %d", result3)
	}
}
