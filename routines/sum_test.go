package routines

import (
	"testing"
)

func TestSumSlices(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8, 9, 10}

	result := SumSlices(slice1, slice2)
	expectedResult := 55

	if result != expectedResult {
		t.Errorf("Expected %d, got %d", expectedResult, result)
	}
}
