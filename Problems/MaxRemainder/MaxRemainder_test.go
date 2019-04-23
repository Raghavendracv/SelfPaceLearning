package MaxRemainder

import "testing"

// TestCaclulateMaxRemainderForValidEvenInputExpectsThree
func TestCaclulateMaxRemainderForValidEvenInputExpectsThree(t *testing.T) {
	actual := CaclulateMaxRemainder([]int{4, 2, 1, 3})
	expected := 3

	if actual != expected {
		t.Errorf("CaclulateMaxRemainder result was incorrect, got: %d, want: %d.", actual, expected)
	}
}

// TestCaclulateMaxRemainderForValidOddInputExpectsFour1
func TestCaclulateMaxRemainderForValidOddInputExpectsFour1(t *testing.T) {
	actual := CaclulateMaxRemainder([]int{2, 1, 5, 3, 4})
	expected := 4

	if actual != expected {
		t.Errorf("CaclulateMaxRemainder result was incorrect, got: %d, want: %d.", actual, expected)
	}
}

// TestCaclulateMaxRemainderForValidOddnputExpectsFour2
func TestCaclulateMaxRemainderForValidOddInputExpectsFour2(t *testing.T) {
	actual := CaclulateMaxRemainder([]int{5, 1, 2, 3, 4})
	expected := 4

	if actual != expected {
		t.Errorf("CaclulateMaxRemainder result was incorrect, got: %d, want: %d.", actual, expected)
	}
}

// TestCaclulateMaxRemainderForValidOddInputExpectsFour3
func TestCaclulateMaxRemainderForValidOddInputExpectsFour3(t *testing.T) {
	actual := CaclulateMaxRemainder([]int{4, 1, 3, 2, 5})
	expected := 4

	if actual != expected {
		t.Errorf("CaclulateMaxRemainder result was incorrect, got: %d, want: %d.", actual, expected)
	}
}

// TestCaclulateMaxRemainderForValidSameElementInputExpectsZero
func TestCaclulateMaxRemainderForValidSameElementInputExpectsZero(t *testing.T) {
	actual := CaclulateMaxRemainder([]int{4, 4, 4, 4, 4})
	expected := 0

	if actual != expected {
		t.Errorf("CaclulateMaxRemainder result was incorrect, got: %d, want: %d.", actual, expected)
	}
}
