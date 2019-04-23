package minelementfrequency

import "testing"

// TestCalculateMinimumElementFrequencyForSameElementSliceOfSizeFourShouldReturnFour
func TestCalculateMinimumElementFrequencyForSameElementSliceOfSizeFourShouldReturnFour(t *testing.T) {
	actual := CalculateMinimumElementFrequency([]int{1, 1, 1, 1})
	expected := 4

	if actual != expected {
		t.Errorf("CalculateMinimumElementFrequency result was incorrect, got: %d, want: %d.", actual, expected)
	}
}

// TestCalculateMinimumElementFrequencyForSameElementSliceOfSizeFourShouldReturnTwo
func TestCalculateMinimumElementFrequencyForSameElementSliceOfSizeFourShouldReturnTwo(t *testing.T) {
	actual := CalculateMinimumElementFrequency([]int{3, 2, 2, 5})
	expected := 2

	if actual != expected {
		t.Errorf("CalculateMinimumElementFrequency result was incorrect, got: %d, want: %d.", actual, expected)
	}
}
