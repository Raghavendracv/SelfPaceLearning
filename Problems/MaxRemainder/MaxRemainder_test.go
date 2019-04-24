package MaxRemainder

import "testing"

// TestCalculateMaxRemainderForValidEvenInputExpectsThree
func TestCalculateMaxRemainderForValidEvenInputExpectsThree(t *testing.T) {
	actual := CalculateMaxRemainder([]int{4, 2, 1, 3})
	expected := 3

	if actual != expected {
		t.Errorf("CalculateMaxRemainder result was incorrect, got: %d, want: %d.", actual, expected)
	}
}

// TestCalculateMaxRemainderForValidOddInputExpectsFour1
func TestCalculateMaxRemainderForValidOddInputExpectsFour1(t *testing.T) {
	actual := CalculateMaxRemainder([]int{2, 1, 5, 3, 4})
	expected := 4

	if actual != expected {
		t.Errorf("CalculateMaxRemainder result was incorrect, got: %d, want: %d.", actual, expected)
	}
}

// TestCalculateMaxRemainderForValidOddInputExpectsFour2
func TestCalculateMaxRemainderForValidOddInputExpectsFour2(t *testing.T) {
	actual := CalculateMaxRemainder([]int{5, 1, 2, 3, 4})
	expected := 4

	if actual != expected {
		t.Errorf("CalculateMaxRemainder result was incorrect, got: %d, want: %d.", actual, expected)
	}
}

// TestCalculateMaxRemainderForValidOddInputExpectsFour3
func TestCalculateMaxRemainderForValidOddInputExpectsFour3(t *testing.T) {
	actual := CalculateMaxRemainder([]int{4, 1, 3, 2, 5})
	expected := 4

	if actual != expected {
		t.Errorf("CalculateMaxRemainder result was incorrect, got: %d, want: %d.", actual, expected)
	}
}

// TestCalculateMaxRemainderForValidSameElementInputExpectsZero
func TestCalculateMaxRemainderForValidSameElementInputExpectsZero(t *testing.T) {
	actual := CalculateMaxRemainder([]int{4, 4, 4, 4, 4})
	expected := 0

	if actual != expected {
		t.Errorf("CalculateMaxRemainder result was incorrect, got: %d, want: %d.", actual, expected)
	}
}
